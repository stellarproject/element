package proxy

func (p *Proxy) Stop() error {
	if p.instance != nil {
		return p.instance.Stop()
	}

	return nil
}
