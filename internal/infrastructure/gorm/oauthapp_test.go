package gorm

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/microserv-io/oauth2-token-vault/internal/domain/models/oauthapp"
)

func TestOAuthAppRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := NewOAuthAppRepository(db)

	testApp := &oauthapp.OAuthApp{
		Provider:     "Test Provider",
		AccessToken:  "test-access-token",
		RefreshToken: "test-refresh-token",
		TokenType:    "Bearer",
		ExpiresAt:    time.Now().Add(1 * time.Hour),
		Scopes:       []string{"scope1", "scope2"},
		OwnerID:      "Test Owner",
	}

	tests := []struct {
		name    string
		setup   func()
		test    func(t *testing.T)
		cleanup func()
	}{
		{
			name: "Create and Find",
			setup: func() {
				if err := repo.Create(context.Background(), testApp); err != nil {
					t.Fatalf("Failed to create OAuthApp: %v", err)
				}
			},
			test: func(t *testing.T) {
				app, err := repo.Find(context.Background(), "Test Owner", "Test Provider")
				if err != nil {
					t.Fatalf("Failed to find OAuthApp: %v", err)
				}
				if app.Provider != testApp.Provider || app.OwnerID != testApp.OwnerID {
					t.Fatalf("Found OAuthApp does not match created one")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM oauth_apps")
			},
		},
		{
			name: "ListForOwner",
			setup: func() {
				if err := repo.Create(context.Background(), testApp); err != nil {
					t.Fatalf("Failed to create OAuthApp: %v", err)
				}
			},
			test: func(t *testing.T) {
				apps, err := repo.ListForOwner(context.Background(), "Test Owner")
				if err != nil {
					t.Fatalf("Failed to list OAuthApps: %v", err)
				}
				log.Printf("apps: %v", apps)
				if len(apps) != 1 || apps[0].OwnerID != "Test Owner" {
					t.Fatalf("ListForOwner returned incorrect results")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM oauth_apps")
			},
		},
		{
			name: "Delete",
			setup: func() {
				if err := repo.Create(context.Background(), testApp); err != nil {
					t.Fatalf("Failed to create OAuthApp: %v", err)
				}
			},
			test: func(t *testing.T) {
				app, err := repo.Find(context.Background(), "Test Owner", "Test Provider")
				if err != nil {
					t.Fatalf("Failed to find OAuthApp: %v", err)
				}

				if err := repo.Delete(context.Background(), app.ID); err != nil {
					t.Fatalf("Failed to delete OAuthApp: %v", err)
				}
				if _, err := repo.Find(context.Background(), "Test Owner", "Test Provider"); err == nil {
					t.Fatalf("Expected error when finding deleted OAuthApp, got none")
				}
			},
			cleanup: func() {
				db.Exec("DELETE FROM oauth_apps")
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
