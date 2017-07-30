package docker

import (
	"github.com/docker/docker/client"
)

const (
	elementRuntimeLabel = "runtime.element"
)

type Docker struct {
	client *client.Client
}

func New() (*Docker, error) {
	c, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Docker{
		client: c,
	}, nil
}
