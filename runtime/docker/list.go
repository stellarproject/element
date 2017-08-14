package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/ehazlett/element/runtime"
)

func (d *Docker) List(namespace string) ([]runtime.Container, error) {
	optFilters := filters.NewArgs()
	optFilters.Add("label", elementRuntimeLabel)

	dockerContainers, err := d.client.ContainerList(context.Background(), types.ContainerListOptions{
		Filters: optFilters,
	})
	if err != nil {
		return nil, err
	}

	var containers []runtime.Container
	for _, c := range dockerContainers {
		endpoint := getContainerEndpoint(c)
		containers = append(containers, Container{
			id:       c.ID,
			endpoint: endpoint,
		})
	}

	return containers, nil
}
