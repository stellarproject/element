package proxy

func (p *Proxy) Update(config *Config) error {
	p.m.Lock()
	defer p.m.Unlock()

	p.config = config

	return nil
}
