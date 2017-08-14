package memory

import (
	"strings"

	"github.com/ehazlett/element/api/types"
)

func (m *Memory) SaveService(service *types.Service) error {
	m.m.Lock()
	m.service[service.ID] = service
	m.m.Unlock()
	return nil
}

func (m *Memory) DeleteService(id string) error {
	m.m.Lock()
	if _, exists := m.service[id]; exists {
		delete(m.service, id)
	}
	m.m.Unlock()
	return nil
}

func (m *Memory) GetServices() ([]*types.Service, error) {
	s := []*types.Service{}
	for _, v := range m.service {
		s = append(s, v)
	}

	return s, nil
}

func (m *Memory) GetServiceByHost(host string) (*types.Service, error) {
	services, err := m.GetServices()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		for _, h := range service.Hosts {
			if strings.Index(h, host) > -1 {
				return service, nil
			}
		}
	}

	return nil, nil
}
