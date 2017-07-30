package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

const (
	defaultListenAddr = ":8080"
	defaultSocketPath = "/var/run/element.sock"
)

// ParseConfig returns a Config object from a raw string config TOML
func ParseConfig(data string) (*Config, error) {
	var cfg Config
	if _, err := toml.Decode(data, &cfg); err != nil {
		return nil, err
	}

	if cfg.ListenAddr == "" {
		logrus.Warnf("using default listen addr: %s", defaultListenAddr)
		cfg.ListenAddr = defaultListenAddr
	}

	if cfg.SocketPath == "" {
		logrus.Warnf("using default socket path: %s", defaultSocketPath)
		cfg.SocketPath = defaultSocketPath
	}

	return &cfg, nil
}
