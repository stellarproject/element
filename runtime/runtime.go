package runtime

import "github.com/ehazlett/element/api/types"

type Container interface {
	ID() string
	Endpoint() string
}

type Runtime interface {
	// Create creates a new container
	Create(spec *types.RuntimeSpec) (Container, error)
	// Delete deletes a container
	Delete(namespace, id string) error
	// List returns all containers in the runtime
	List(namespace string) ([]Container, error)
	// Get returns the specified container
	Get(namespace, id string) (Container, error)
}
