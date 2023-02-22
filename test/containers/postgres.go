package containers

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Drozd0f/ttto-go/pkg/config"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	user     string
	password string
	name     string
	instance testcontainers.Container
}

func NewTestDatabase(t *testing.T, c config.DbConfig) (*TestDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	port, err := nat.NewPort("tcp", c.Port)
	if err != nil {
		return nil, fmt.Errorf("nat new port: %w", err)
	}
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14",
		ExposedPorts: []string{fmt.Sprintf("%s/tcp", c.Port)},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_USER":     c.User,
			"POSTGRES_PASSWORD": c.Password,
			"POSTGRES_DB":       c.Name,
		},
		WaitingFor: wait.ForListeningPort(port),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	return &TestDatabase{
		instance: postgres,
	}, nil
}

func (db *TestDatabase) Port(t *testing.T) int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.instance.MappedPort(ctx, "5432")
	require.NoError(t, err)
	return p.Int()
}

func (db *TestDatabase) ConnectionString(t *testing.T) string {
	return fmt.Sprintf("postgres://%s:%s@127.0.0.1:%d/%s?sslmode=disable", db.user, db.password, db.Port(t), db.name)
}

func (db *TestDatabase) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.instance.Terminate(ctx))
}
