package memory

import (
	"sync"

	"github.com/ehazlett/element/api/types"
	"github.com/ehazlett/element/proxy"
)

type Memory struct {
	proxy   map[string]*proxy.Config
	service map[string]*types.Service
	m       sync.Mutex
}

func NewMemory() (*Memory, error) {
	return &Memory{
		proxy:   map[string]*proxy.Config{},
		service: map[string]*types.Service{},
		m:       sync.Mutex{},
	}, nil
}
