package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfig(t *testing.T) {
	wd, _ := os.Getwd()
	rootDir := filepath.Join(wd, "../..")
	cfgPath := filepath.Join(rootDir, "tests/testdata/config")

	tests := []struct {
		name       string
		cfgFile    string
		envVars    map[string]string
		wantErr    bool
		wantConfig *Config
	}{
		{
			name:    "valid config file",
			cfgFile: "valid_config",
			envVars: map[string]string{
				"PROVIDER__0__CLIENT_SECRET": "env_client_secret",
			},
			wantErr: false,
			wantConfig: &Config{
				Providers: []Provider{
					{
						Name:         "google",
						ClientID:     "google-client-id",
						ClientSecret: "env_client_secret",
						RedirectURL:  "http://localhost:8080/oauth/google/callback",
						AuthURL:      "https://accounts.google.com/o/oauth2/auth",
						TokenURL:     "https://accounts.google.com/o/oauth2/token",
						Scopes:       []string{"https://www.googleapis.com/auth/analytics"},
					},
				},
				AllowProviderRegistration: true,
			},
		},
		{
			name:       "valid without client secret",
			cfgFile:    "valid_config",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing client id",
			cfgFile:    "missing_client_id",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing provider name",
			cfgFile:    "missing_provider_name",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing redirect url",
			cfgFile:    "missing_redirect_url",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing auth url",
			cfgFile:    "missing_auth_url",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing token url",
			cfgFile:    "missing_token_url",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
		{
			name:       "missing scopes",
			cfgFile:    "missing_scopes",
			envVars:    map[string]string{},
			wantErr:    true,
			wantConfig: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			gotConfig, err := NewConfig(cfgPath, tt.cfgFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareConfigs(gotConfig, tt.wantConfig) {
				t.Errorf("NewConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}

func (p Provider) NotEqual(other Provider) bool {
	if p.Name != other.Name {
		return true
	}
	if p.ClientID != other.ClientID {
		return true
	}
	if p.ClientSecret != other.ClientSecret {
		return true
	}
	if p.RedirectURL != other.RedirectURL {
		return true
	}
	if p.AuthURL != other.AuthURL {
		return true
	}
	if p.TokenURL != other.TokenURL {
		return true
	}
	if len(p.Scopes) != len(other.Scopes) {
		return true
	}
	for i := range p.Scopes {
		if p.Scopes[i] != other.Scopes[i] {
			return true
		}
	}
	return false
}

func compareConfigs(a, b *Config) bool {
	if a.AllowProviderRegistration != b.AllowProviderRegistration {
		return false
	}
	if len(a.Providers) != len(b.Providers) {
		return false
	}
	for i := range a.Providers {
		if a.Providers[i].NotEqual(b.Providers[i]) {
			return false
		}
	}
	return true
}
