package config

// Config is the top level application configuration
type Config struct {
	ListenAddr string
	GRPCAddr   string
	SocketPath string
	Runtime    *Runtime
	Datastore  string
}

type Runtime struct {
	Name   string
	Config interface{}
}
