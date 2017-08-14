package docker

import "github.com/ehazlett/element/runtime"

func (d *Docker) Get(namespace, id string) (runtime.Container, error) {
	containers, err := d.List(namespace)
	if err != nil {
		return nil, err
	}

	for _, c := range containers {
		if c.ID() == id {
			return c, nil
		}
	}

	return nil, ErrContainerNotFound
}
