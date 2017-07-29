package server

import (
	"net/http"

	"github.com/ehazlett/element/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	return &Server{
		cfg: cfg,
	}, nil
}

func (s *Server) Run() error {
	if s.cfg.EnableMetrics {
		// start prometheus listener
		http.Handle("/metrics", prometheus.Handler())
		go func() {
			if err := http.ListenAndServe(s.cfg.ListenAddr, nil); err != nil {
				logrus.Error("unable to start metric listener: %s", err)
			}
		}()
	}

	r, err := s.router()
	if err != nil {
		return err
	}

	http.Handle("/", r)

	if err := http.ListenAndServe(s.cfg.ListenAddr, nil); err != nil {
		return err
	}

	return nil
}
