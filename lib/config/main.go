package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

// Config variables for grafana.com.
type Config struct {
	address string
	apiKey string
	tenantID string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getVar(varName string) string {
	valFromEnv := viper.GetString(varName)

	if strings.HasPrefix(valFromEnv, "/") {
		dat, err := ioutil.ReadFile(valFromEnv)
		check(err)

		return string(dat)
	}

	return valFromEnv

}

// Init Set the environment Prefix
func (c *Config) Init() {
	viper.SetEnvPrefix("cortex")
	viper.BindEnv("tenant_id")
	viper.BindEnv("address")
	viper.BindEnv("api_key")

	c.address = getVar("ADDRESS")

	// Loop over the string with len.
	fmt.Printf("Address is: %s", c.address)
	fmt.Print("\n")

	c.apiKey = getVar("API_KEY")

	// Loop over the string with len.
	fmt.Print("API Key is: ")
	for i := 0; i < len(c.apiKey); i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")

	c.tenantID = getVar("TENANT_ID")

	// Loop over the string with len.
	fmt.Print("Tenant ID is:")
	for i := 0; i < len(c.tenantID); i++ {
		fmt.Print("*")
	}

	fmt.Print("\n")

}

// Address of the grafana.com API Endpoint
func (c *Config) Address() string {
	return c.address
}

// APIKey from grafana.com
func (c *Config) APIKey() string {
	return c.apiKey
}

// TenantID from grafana.com
func (c *Config) TenantID() string {
	return c.tenantID
}
