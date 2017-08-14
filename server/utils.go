package server

import (
	"errors"
	"net/url"

	"github.com/ehazlett/element/config"
	"github.com/ehazlett/element/datastore"
	"github.com/ehazlett/element/datastore/memory"
	"github.com/ehazlett/element/runtime"
	"github.com/ehazlett/element/runtime/docker"
)

var (
	ErrInvalidRuntime   = errors.New("invalid runtime specified")
	ErrInvalidDatastore = errors.New("invalid datastore specified")
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

// loadDatastore loads a datastore from the specified configuration
func loadDatastore(ds string) (datastore.Datastore, error) {
	u, err := url.Parse(ds)
	if err != nil {
		return nil, err
	}

	var d datastore.Datastore
	switch u.Scheme {
	case "memory":
		m, err := memory.NewMemory()
		if err != nil {
			return nil, err
		}
		d = m
	default:
		return nil, ErrInvalidDatastore
	}

	return d, nil
}
