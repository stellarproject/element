package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ehazlett/element/config"
	"github.com/ehazlett/element/proxy"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg   *config.Config
	proxy *proxy.Proxy
}

func NewServer(cfg *config.Config) (*Server, error) {
	p, err := proxy.NewProxy(&proxy.Config{})
	if err != nil {
		return nil, err
	}
	return &Server{
		cfg:   cfg,
		proxy: p,
	}, nil
}

func (s *Server) Run() error {
	r := s.router()

	srv := &http.Server{
		Handler: r,
	}

	go func() {
		// check for existing socket
		if _, err := os.Stat(s.cfg.SocketPath); err == nil {
			os.Remove(s.cfg.SocketPath)
		}
		l, err := net.Listen("unix", s.cfg.SocketPath)
		if err != nil {
			logrus.Errorf("unable to start element server: %s", err)
			return
		}

		srv.Serve(l)
	}()

	cfg := &proxy.Config{
		Frontends: map[string]*proxy.Frontend{
			"element": &proxy.Frontend{
				Name:  "element",
				Hosts: []string{s.cfg.ListenAddr},
				Backend: &proxy.Backend{
					Path:      "/",
					Upstreams: []string{"unix:" + s.cfg.SocketPath},
				},
			},
		},
	}

	if err := s.proxy.Update(cfg); err != nil {
		return err
	}

	if err := s.proxy.Start(); err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for range c {
			if err := s.proxy.Reload(); err != nil {
				logrus.Errorf("error reloading proxy: %s", err)
			}
		}
	}()

	s.proxy.Wait()

	return nil
}

func (s *Server) Stop() error {
	return s.proxy.Stop()
}
