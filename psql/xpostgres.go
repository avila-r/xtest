package psql

import (
	"context"
	"log"
	"time"

	"github.com/avila-r/g"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type XPostgresContainer struct {
	// Just don't initialize it
	Container *postgres.PostgresContainer

	// We recommend initializing that with
	// context.Background (new empty context)
	Context context.Context
}

type Options struct {
	DatabaseName string
	Username     string
	Password     string
}

var (
	DefaultOptions = &Options{
		DatabaseName: "test-db",
		Username:     "root",
		Password:     "",
	}
)

func (p *XPostgresContainer) Start(options *Options) {
	var (
		db_name  = g.If(options.DatabaseName != "", options.DatabaseName, "test-db")
		username = g.If(options.Username != "", options.Username, "test-user")
		password = options.Password
	)

	container, err := postgres.Run(
		p.Context,
		"postgres:16.1",
		postgres.WithDatabase(db_name),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	if err != nil {
		log.Printf("failed to start container - %s", err)
	}

	p.Container = container
}

func (p *XPostgresContainer) Terminate() {
	if err := p.Container.Terminate(p.Context); err != nil {
		log.Printf("failed to terminate container - %s", err)
	}
}

func (p *XPostgresContainer) StartOrError(options *Options) error {
	var (
		db_name  = g.If(options.DatabaseName != "", options.DatabaseName, "test-db")
		username = g.If(options.Username != "", options.Username, "test-user")
		password = options.Password
	)

	container, err := postgres.Run(
		p.Context,
		"postgres:16.1",
		postgres.WithDatabase(db_name),
		postgres.WithUsername(username),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)

	if err != nil {
		return err
	}

	p.Container = container
	return nil
}

func (p *XPostgresContainer) TerminateOrError() error {
	if err := p.Container.Terminate(p.Context); err != nil {
		return err
	}

	return nil
}
