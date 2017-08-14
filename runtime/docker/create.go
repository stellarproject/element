package docker

import (
	"fmt"
	"strconv"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/go-connections/nat"
	"github.com/ehazlett/element/api/types"
	"github.com/ehazlett/element/runtime"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func (d *Docker) Create(spec *types.RuntimeSpec) (runtime.Container, error) {
	if spec.Labels == nil {
		spec.Labels = map[string]string{}
	}
	labels := spec.Labels
	// insert element labels
	labels[elementRuntimeLabel] = "docker"

	p := strconv.Itoa(int(spec.Port))
	port, err := nat.NewPort(spec.Protocol, p)
	if err != nil {
		return nil, err
	}
	containerConfig := &container.Config{
		Image:  spec.Image,
		Labels: spec.Labels,
		ExposedPorts: nat.PortSet{
			port: struct{}{},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			port: []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
				},
			},
		},
	}

	logrus.Debugf("%+v", containerConfig)

	// create
	resp, err := d.client.ContainerCreate(context.Background(), containerConfig, hostConfig, nil, "")
	if err != nil {
		return nil, err
	}

	// start
	if err := d.client.ContainerStart(context.Background(), resp.ID, dockertypes.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	optFilters := filters.NewArgs()
	optFilters.Add("id", resp.ID)
	optFilters.Add("label", elementRuntimeLabel)

	containers, err := d.client.ContainerList(context.Background(), dockertypes.ContainerListOptions{
		Filters: optFilters,
	})
	if err != nil {
		return nil, err
	}

	if len(containers) == 0 {
		return nil, fmt.Errorf("error creating container: no container found after start")
	}

	container := containers[0]
	endpoint := getContainerEndpoint(container)

	return Container{
		id:       container.ID,
		endpoint: endpoint,
	}, nil
}
