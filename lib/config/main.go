package config

import (
	"github.com/spf13/viper"
)

// Config variables for grafana.com.
type Config struct {
}

// Init Set the environment Prefix
func (c *Config) Init() {
	viper.SetEnvPrefix("cortex")
	viper.BindEnv("tenant_id")
	viper.BindEnv("address")
	viper.BindEnv("api_key")
}

// Address of the grafana.com API Endpoint
func (c *Config) Address() string {
	return viper.GetString("ADDRESS")
}

// APIKey from grafana.com
func (c *Config) APIKey() string {
	return viper.GetString("API_KEY")
}

// TenantID from grafana.com
func (c *Config) TenantID() string {
	return viper.GetString("TENANT_ID")
}
