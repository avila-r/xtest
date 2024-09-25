package psql_test

import (
	"testing"

	"github.com/avila-r/xtest"
	"github.com/avila-r/xtest/psql"
	"github.com/stretchr/testify/assert"
)

var (
	options = &psql.Options{
		DatabaseName: "test-db",
		Username:     "root",
		Password:     "123",
	}
)

func Test_Container(t *testing.T) {
	assert := assert.New(t)

	c := xtest.PostgresContainer

	t.Cleanup(func() {
		c.Terminate()
	})

	err := c.StartOrError(options)

	assert.Nil(err)
}
