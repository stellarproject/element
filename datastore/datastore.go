package datastore

import (
	"github.com/ehazlett/element/api/types"
	"github.com/ehazlett/element/proxy"
)

type Datastore interface {
	// proxy
	SaveProxyConfig(id string, cfg *proxy.Config) error
	DeleteProxyConfig(id string) error
	GetProxyConfigs() ([]*proxy.Config, error)
	// services
	SaveService(service *types.Service) error
	DeleteService(id string) error
	GetServices() ([]*types.Service, error)
	GetServiceByHost(host string) (*types.Service, error)
}
