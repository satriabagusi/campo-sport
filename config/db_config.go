package config

import "github.com/satriabagusi/campo-sport/pkg/utility"

type Config struct {
	PostgresConnectionString string
	ServerAddress            string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.PostgresConnectionString = utility.GetEnv("CONNECTION_STRING")
	c.ServerAddress = utility.GetEnv("SERVER_ADDRESS")
}
