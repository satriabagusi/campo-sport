package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

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
	midtrans.ServerKey = utility.GetEnv("MIDTRANS_SERVER_KEY")
	midtrans.Environment = midtrans.Sandbox
}
