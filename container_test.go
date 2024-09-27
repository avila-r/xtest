package xtest_test

import (
	"testing"

	"github.com/avila-r/xtest"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_GenericContainer(t *testing.T) {
	assert := assert.New(t)

	c, err := xtest.NewContainer(&xtest.ContainerRequest{
		ContainerName: "go-test-redis",
		Image:         "redis",
		ExposedPorts:  []string{"6379/tcp"},
		WaitingFor:    wait.ForLog("Ready to accept connections"),
	})

	assert.Nil(err)

	t.Cleanup(func() {
		c.Terminate()
	})

	c.Start()

	if running := c.Advanced.IsRunning(); !running {
		t.Errorf("container isnt running")
	}
}
