package docker

type Container struct {
	id string
}

func (c Container) ID() string {
	return c.id
}
