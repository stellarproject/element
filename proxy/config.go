package proxy

import (
	"bytes"
	"text/template"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Frontends map[string]*Frontend `json:"frontends,omitempty"`
}

type Frontend struct {
	Name    string   `json:"name"`
	Hosts   []string `json:"hosts,omitempty"`
	Backend *Backend `json:"backend,omitempty"`
}

type Backend struct {
	Path      string   `json:"path"`
	Upstreams []string `json:"upstreams,omitempty"`
}

func (c *Config) Body() []byte {
	t := template.New("proxy")
	tmpl, err := t.Parse(configTemplate)
	if err != nil {
		logrus.Errorf("error parsing proxy template: %s", err)
		return nil
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, c); err != nil {
		logrus.Errorf("error executing proxy template: %s", err)
		return nil
	}

	return b.Bytes()
}

func (c *Config) Path() string {
	return ""
}

func (c *Config) ServerType() string {
	return "http"
}
