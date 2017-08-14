package server

import (
	configurationapi "github.com/ehazlett/element/api/services/configuration"
	"golang.org/x/net/context"
)

func (s *Server) CreateService(ctx context.Context, req *configurationapi.CreateServiceRequest) (*configurationapi.CreateServiceResponse, error) {
	// TODO
	// save to datastore
	if err := s.store.SaveService(req.Service); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Server) ListServices(ctx context.Context, req *configurationapi.ListServicesRequest) (*configurationapi.ListServicesResponse, error) {
	// TODO
	return nil, nil
}
