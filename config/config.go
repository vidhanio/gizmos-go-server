package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerHost string
	MongoPort  string
}

func (config *Config) initialize() {
	config.ServerHost = os.Getenv("SERVER_PORT")
	config.MongoPort = os.Getenv("MONGO_PORT")

	if config.ServerHost == "" {
		config.ServerHost = "localhost:8000"
	}

	if config.MongoPort == "" {
		config.MongoPort = "27017"
	}
}

func (config *Config) MongoURI() string {
	return fmt.Sprintf(
		"mongodb://localhost:%s",
		config.MongoPort,
	)
}

func New() *Config {
	config := new(Config)
	config.initialize()
	return config
}
