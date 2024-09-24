package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Provider struct {
	Name         string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
	TokenURL     string
	Scopes       []string
}

type Config struct {
	Providers                 []Provider
	AllowProviderRegistration bool
}

func NewConfig(cfgPath string) (*Config, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath(cfgPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var configObject Config
	if err := viper.Unmarshal(&configObject); err != nil {
		return nil, err
	}

	return &configObject, nil
}
