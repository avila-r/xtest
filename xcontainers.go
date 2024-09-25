package xtest

import (
	"context"

	"github.com/avila-r/xtest/psql"
)

var (
	PostgresContainer = func() *psql.XPostgresContainer {
		return &psql.XPostgresContainer{
			Context: context.Background(),
		}
	}()
)
