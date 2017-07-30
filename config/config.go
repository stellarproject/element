package config

// Config is the top level configuration
type Config struct {
	ListenAddr string
	SocketPath string
	Runtime    *Runtime
}

type Runtime struct {
	Name   string
	Config interface{}
}
