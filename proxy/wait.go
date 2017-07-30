package proxy

func (p *Proxy) Wait() {
	if p.instance != nil {
		p.instance.Wait()
	}
}
