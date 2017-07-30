package proxy

func (p *Proxy) Reload() error {
	p.m.Lock()
	defer p.m.Unlock()

	i, err := p.instance.Restart(p.config)
	if err != nil {
		return err
	}

	p.instance = i

	return nil
}
