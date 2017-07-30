package proxy

import (
	"errors"
	"sync"

	"github.com/ehazlett/element/version"
	"github.com/mholt/caddy"
	_ "github.com/mholt/caddy/caddyhttp"
)

var (
	ErrFrontendExists       = errors.New("frontend exists")
	ErrFrontendDoesNotExist = errors.New("frontend does not exist")
)

type Proxy struct {
	config   *Config
	instance *caddy.Instance
	m        sync.Mutex
}

func NewProxy(config *Config) (*Proxy, error) {
	if config.Frontends == nil {
		config.Frontends = map[string]*Frontend{}
	}
	caddy.AppName = "element"
	caddy.AppVersion = version.Version + version.Build

	return &Proxy{
		config: config,
		m:      sync.Mutex{},
	}, nil
}

func (p *Proxy) Config() (*Config, error) {
	return p.config, nil
}
