package runtime

type Container interface {
	ID() string
}

type Runtime interface {
	// Create creates a new container
	Create(spec *Spec) error
	// Delete deletes a container
	Delete(id string) error
	// List returns all containers in the runtime
	List(namespace string) ([]Container, error)
}
