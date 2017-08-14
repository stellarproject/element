package memory

import "github.com/ehazlett/element/proxy"

func (m *Memory) SaveProxyConfig(id string, cfg *proxy.Config) error {
	m.m.Lock()
	m.proxy[id] = cfg
	m.m.Unlock()
	return nil
}

func (m *Memory) DeleteProxyConfig(id string) error {
	m.m.Lock()
	if _, exists := m.proxy[id]; exists {
		delete(m.proxy, id)
	}
	m.m.Unlock()
	return nil
}

func (m *Memory) GetProxyConfigs() ([]*proxy.Config, error) {
	c := []*proxy.Config{}
	for _, v := range m.proxy {
		c = append(c, v)
	}

	return c, nil
}
