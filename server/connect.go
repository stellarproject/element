package server

import (
	"github.com/ehazlett/element/proxy"
	"github.com/sirupsen/logrus"
)

func (s *Server) connect(host string) error {
	service, err := s.store.GetServiceByHost(host)
	if err != nil {
		return err
	}

	if service == nil {
		return ErrServiceNotFound
	}

	// create container
	container, err := s.runtime.Create(service.RuntimeSpec)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"id":       container.ID(),
		"endpoint": container.Endpoint(),
	}).Debug("container created")

	// configure proxy
	frontend := &proxy.Frontend{
		Name:  service.ID,
		Hosts: service.Hosts,
		Backend: &proxy.Backend{
			Path:      "/",
			Upstreams: []string{container.Endpoint()},
		},
	}
	logrus.WithFields(logrus.Fields{
		"frontend": frontend,
	}).Debug("configuring proxy")

	if err := s.proxy.AddFrontend(frontend); err != nil {
		return err
	}

	// reload
	if err := s.proxy.Reload(); err != nil {
		return err
	}

	return nil
}
