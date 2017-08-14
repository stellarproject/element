package docker

type Container struct {
	id       string
	endpoint string
}

func (c Container) ID() string {
	return c.id
}

func (c Container) Endpoint() string {
	// TODO
	return c.endpoint
}
