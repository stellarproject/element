package server

import (
	"errors"

	"github.com/ehazlett/element/config"
	"github.com/ehazlett/element/runtime"
	"github.com/ehazlett/element/runtime/docker"
)

var (
	ErrInvalidRuntime = errors.New("invalid runtime specified")
)

// loadRuntime loads a runtime from the specified configuration
func loadRuntime(cfg *config.Runtime) (runtime.Runtime, error) {
	var rt runtime.Runtime

	switch cfg.Name {
	case "docker":
		r, err := docker.New()
		if err != nil {
			return nil, err
		}
		rt = r
	default:
		return nil, ErrInvalidRuntime
	}

	return rt, nil
}
