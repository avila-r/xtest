package xtest

import (
	"context"

	"github.com/avila-r/g"
	d "github.com/testcontainers/testcontainers-go"
)

type Container struct {
	Advanced d.Container
	Context  context.Context
}

func (c *Container) Start() {
	c.Advanced.Start(c.Context)
}

func (c *Container) Terminate() {
	c.Advanced.Terminate(c.Context)
}

func (c *Container) Endpoint(port ...string) (string, error) {
	var p string

	if len(port) == 0 {
		p = ""
	} else {
		p = g.Coalesce(port...)
	}

	return c.Advanced.Endpoint(c.Context, p)
}
