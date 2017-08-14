package server

import (
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	configurationapi "github.com/ehazlett/element/api/services/configuration"
	"github.com/ehazlett/element/config"
	"github.com/ehazlett/element/datastore"
	"github.com/ehazlett/element/proxy"
	"github.com/ehazlett/element/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	ErrServiceNotFound = errors.New("service not found")
)

type Server struct {
	cfg     *config.Config
	proxy   *proxy.Proxy
	runtime runtime.Runtime
	store   datastore.Datastore
}

func NewServer(cfg *config.Config) (*Server, error) {
	p, err := proxy.NewProxy(&proxy.Config{})
	if err != nil {
		return nil, err
	}
	r, err := loadRuntime(cfg.Runtime)
	if err != nil {
		return nil, err
	}

	store, err := loadDatastore(cfg.Datastore)
	if err != nil {
		return nil, err
	}

	c, err := r.List("")
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"containers": c,
	}).Debug("runtime containers")

	return &Server{
		cfg:     cfg,
		proxy:   p,
		runtime: r,
		store:   store,
	}, nil
}

func (s *Server) Run() error {
	grpcServer := grpc.NewServer()
	configurationapi.RegisterConfigurationServer(grpcServer, s)

	l, err := net.Listen("tcp", s.cfg.GRPCAddr)
	if err != nil {
		return err
	}

	//cfg := &proxy.Config{
	//	Frontends: map[string]*proxy.Frontend{
	//		"element": &proxy.Frontend{
	//			Name:  "element",
	//			Hosts: []string{s.cfg.ListenAddr},
	//			Backend: &proxy.Backend{
	//				Path:      "/",
	//				Upstreams: []string{"unix:" + s.cfg.SocketPath},
	//			},
	//		},
	//	},
	//}

	if err := s.proxy.Start(); err != nil {
		return err
	}

	// handle SIGHUP
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			logrus.Debugf("received SIGHUP; reloading")
			if err := s.proxy.Reload(); err != nil {
				logrus.Errorf("error reloading proxy: %s", err)
			}
		}
	}()

	grpcServer.Serve(l)

	return nil
}

func (s *Server) Stop() error {
	return s.proxy.Stop()
}
