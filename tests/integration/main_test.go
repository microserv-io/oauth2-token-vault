//go:build integration

package integration

import (
	"database/sql"
	"fmt"
	"github.com/microserv-io/oauth-credentials-server/internal"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

const ServerPort = "50052"

func setupContainers() (func(), error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Panicf("failed to create docker pool: %v", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Panicf("failed to ping docker: %v", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "17",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")

	if err := os.Setenv("DATABASE_HOST", "localhost"); err != nil {
		return nil, fmt.Errorf("failed to set DATABASE_HOST env variable: %v", err)
	}
	if err := os.Setenv("DATABASE_PORT", strings.Split(hostAndPort, ":")[1]); err != nil {
		return nil, fmt.Errorf("failed to set DATABASE_PORT env variable: %v", err)
	}

	if err := os.Setenv("DATABASE_USER", "user_name"); err != nil {
		return nil, fmt.Errorf("failed to set DATABASE_USER env variable: %v", err)
	}

	if err := os.Setenv("DATABASE_PASSWORD", "secret"); err != nil {
		return nil, fmt.Errorf("failed to set DATABASE_PASSWORD env variable: %v", err)
	}

	if err := os.Setenv("DATABASE_NAME", "dbname"); err != nil {
		return nil, fmt.Errorf("failed to set DATABASE_NAME env variable: %v", err)
	}

	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	if err := resource.Expire(120); err != nil {
		return nil, err
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}, nil
}

func startServer() {

	application, err := internal.NewApplication("config")
	if err != nil {
		log.Panicf("failed to create application: %v", err)
	}

	go func() {
		if err := application.Run("50052"); err != nil {
			log.Panicf("failed to run application: %v", err)
		}
		defer application.Stop()
	}()

	// Wait for the server to start
	time.Sleep(2 * time.Second)
}

func TestMain(m *testing.M) {

	purgeFunc, err := setupContainers()
	if err != nil {
		log.Fatalf("Could not set up containers: %v", err)
	}
	defer purgeFunc()

	startServer()

	m.Run()

}
