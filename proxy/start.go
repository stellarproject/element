package proxy

import "github.com/mholt/caddy"

func (p *Proxy) Start() error {
	p.m.Lock()
	defer p.m.Unlock()

	i, err := caddy.Start(p.config)
	if err != nil {
		return err
	}

	p.instance = i

	return nil
}
