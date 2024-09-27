package xtest

import (
	"context"

	d "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type ContainerRequest struct {
	ContainerName string
	Image         string
	ExposedPorts  []string
	Env           map[string]string
	Cmd           []string
	Labels        map[string]string
	WaitingFor    wait.Strategy
}

func NewContainer(c *ContainerRequest) (*Container, error) {
	ctx := context.Background()

	request := d.ContainerRequest{
		Name:         c.ContainerName,
		Image:        c.Image,
		ExposedPorts: c.ExposedPorts,
		Env:          c.Env,
		Cmd:          c.Cmd,
		Labels:       c.Labels,
		WaitingFor:   c.WaitingFor,
	}

	container, err := d.GenericContainer(ctx, d.GenericContainerRequest{
		ContainerRequest: request,
		Started:          false,
	})

	if err != nil {
		return nil, err
	}

	response := &Container{
		Advanced: container,
		Context:  ctx,
	}

	return response, nil
}
