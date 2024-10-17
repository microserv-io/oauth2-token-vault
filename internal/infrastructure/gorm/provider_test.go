package gorm

import (
	"context"
	"testing"
	"time"

	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/provider"
)

func TestProviderRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewProviderRepository(db)

	testProvider := &provider.Provider{
		Name:         "Test Provider",
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost/callback",
		AuthURL:      "http://localhost/auth",
		TokenURL:     "http://localhost/token",
		Scopes:       []string{"scope1", "scope2"},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	tests := []struct {
		name    string
		setup   func()
		test    func(t *testing.T)
		cleanup func()
	}{
		{
			name: "Create and FindByName",
			setup: func() {
				if err := repo.Create(context.Background(), testProvider); err != nil {
					t.Fatalf("Failed to create Provider: %v", err)
				}
			},
			test: func(t *testing.T) {
				providerObj, err := repo.FindByName(context.Background(), "Test Provider")
				if err != nil {
					t.Fatalf("Failed to find Provider: %v", err)
				}
				if providerObj.Name != testProvider.Name {
					t.Fatalf("Found Provider does not match created one")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM providers")
			},
		},
		{
			name: "List",
			setup: func() {
				if err := repo.Create(context.Background(), testProvider); err != nil {
					t.Fatalf("Failed to create Provider: %v", err)
				}
			},
			test: func(t *testing.T) {
				providers, err := repo.List(context.Background())
				if err != nil {
					t.Fatalf("Failed to list Providers: %v", err)
				}
				if len(providers) != 1 || providers[0].Name != "Test Provider" {
					t.Fatalf("List returned incorrect results")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM providers")
			},
		},
		{
			name: "Delete",
			setup: func() {
				if err := repo.Create(context.Background(), testProvider); err != nil {
					t.Fatalf("Failed to create Provider: %v", err)
				}
			},
			test: func(t *testing.T) {
				if err := repo.Delete(context.Background(), testProvider.Name); err != nil {
					t.Fatalf("Failed to delete Provider: %v", err)
				}
				_, err := repo.FindByName(context.Background(), "Test Provider")
				if err == nil {
					t.Fatalf("Expected error when finding deleted Provider, got none")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM providers")
			},
		},
		{
			name: "Update",
			setup: func() {
				if err := repo.Create(context.Background(), testProvider); err != nil {
					t.Fatalf("Failed to create Provider: %v", err)
				}
			},
			test: func(t *testing.T) {
				providerObj, err := repo.FindByName(context.Background(), "Test Provider")
				if err != nil {
					t.Fatalf("Failed to find Provider: %v", err)
				}
				providerObj.ClientID = "updated-client-id"
				if err := repo.Update(context.Background(), providerObj); err != nil {
					t.Fatalf("Failed to update Provider: %v", err)
				}
				updatedProvider, err := repo.FindByName(context.Background(), "Test Provider")
				if err != nil {
					t.Fatalf("Failed to find updated Provider: %v", err)
				}
				if updatedProvider.ClientID != "updated-client-id" {
					t.Fatalf("Provider was not updated correctly")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM providers")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			tt.test(t)
			tt.cleanup()
		})
	}
}
