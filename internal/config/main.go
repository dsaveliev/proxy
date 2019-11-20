package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var once sync.Once
var cfg *Config

// Config struct declares all possible configuration variables
type Config struct {
	Addr             string `envconfig:"ADDR"`
	ProxyEndpoint    string `envconfig:"PROXY_ENDPOINT"`
	CertPath         string `envconfig:"CERT_PATH"`
	ConcurrencyLimit int    `envconfig:"CONCURRENCY_LIMIT"`
	MaxIdleConns     int    `envconfig:"MAX_IDLE_CONNECTIONS"`
	IdleConnTimeout  int    `envconfig:"IDLE_CONN_TIMEOUT"`
	ServerTimeout    int    `envconfig:"SERVER_TIMEOUT"`
	ClientTimeout    int    `envconfig:"CLIENT_TIMEOUT"`
	DefaultTopValue  int    `envconfig:"DEFAULT_TOP_VALUE"`
	DefaultSkipValue int    `envconfig:"DEFAULT_SKIP_VALUE"`
}

// Godotenv leverage environment variables to configure application
func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}

	cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal("Error load env variables: " + err.Error())
	}
}

// SetConfig initialize cfg for the purposes of testing
func SetConfig(c *Config) {
	cfg = c
}

// GetConfig method reads, initializes and returns the proxy config
func GetConfig() *Config {
	once.Do(initialize)
	return cfg
}
