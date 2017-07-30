package proxy

import "github.com/sirupsen/logrus"

func (p *Proxy) AddFrontend(f *Frontend) error {
	p.m.Lock()
	defer p.m.Unlock()

	if _, ok := p.config.Frontends[f.Name]; ok {
		return ErrFrontendExists
	}

	p.config.Frontends[f.Name] = f
	logrus.WithFields(logrus.Fields{
		"name":  f.Name,
		"hosts": f.Hosts,
	}).Debug("frontend added")
	return nil
}

func (p *Proxy) RemoveFrontend(name string) error {
	p.m.Lock()
	defer p.m.Unlock()

	if _, ok := p.config.Frontends[name]; ok {
		delete(p.config.Frontends, name)
		logrus.WithFields(logrus.Fields{
			"name": name,
		}).Debug("frontend removed")
	}

	return nil
}

func (p *Proxy) UpdateFrontend(f *Frontend) error {
	p.m.Lock()
	defer p.m.Unlock()

	if _, ok := p.config.Frontends[f.Name]; !ok {
		return ErrFrontendDoesNotExist
	}

	p.config.Frontends[f.Name] = f
	logrus.WithFields(logrus.Fields{
		"name": f.Name,
	}).Debug("frontend updated")

	return nil
}
