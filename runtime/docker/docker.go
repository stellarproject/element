package docker

import (
	"errors"

	"github.com/docker/docker/client"
)

const (
	elementRuntimeLabel = "element.runtime"
)

var (
	ErrContainerNotFound = errors.New("container not found")
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
