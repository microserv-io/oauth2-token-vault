package tests

import (
	"github.com/microserv-io/oauth-credentials-server/internal"
	"github.com/microserv-io/oauth-credentials-server/internal/infrastructure/gorm"
	"log"
	"strings"
)

import "github.com/ory/dockertest/v3"

type TestApp struct {
	*internal.Application

	pool      *dockertest.Pool
	resources []*dockertest.Resource

	Port         int
	Repositories struct {
		OAuthApp *gorm.OAuthAppRepository
		Provider *gorm.ProviderRepository
	}
}

func NewTestApp(configPath string) *TestApp {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	databaseResource, err := setupDatabase(pool)
	if err != nil {
		log.Panicf("failed to setup database: %v", err)
	}

	log.Print(databaseResource.GetHostPort("5432/tcp"))

	db, err := gorm.Open(
		"host=localhost user=user_name password=secret dbname=dbname port="+strings.Split(databaseResource.GetHostPort("5432/tcp"), ":")[1]+" sslmode=disable",
		true,
		20,
		1,
		nil,
	)
	if err != nil {
		log.Panicf("failed to open database: %v", err)
	}

	application, err := internal.NewApplication(
		configPath,
		internal.WithDatabase(db),
	)
	if err != nil {
		log.Panicf("failed to create application: %v", err)
	}

	go func() {
		if err := application.Run(); err != nil {
			log.Panicf("failed to run application: %v", err)
		}
	}()

	gorm.NewOAuthAppRepository(db)
	gorm.NewProviderRepository(db)

	return &TestApp{
		Application: application,

		pool:      pool,
		resources: []*dockertest.Resource{databaseResource},

		Repositories: struct {
			OAuthApp *gorm.OAuthAppRepository
			Provider *gorm.ProviderRepository
		}{
			OAuthApp: gorm.NewOAuthAppRepository(db),
			Provider: gorm.NewProviderRepository(db),
		},
		Port: application.GetPort(),
	}
}

func (app *TestApp) Cleanup() {
	for _, resource := range app.resources {
		if err := app.pool.Purge(resource); err != nil {
			log.Printf("failed to purge resource: %v", err)
		}
	}
	app.Stop()
}
