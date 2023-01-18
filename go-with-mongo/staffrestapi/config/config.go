package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server   string
	Database string
}

// Parse the config file 'config.toml', and establish a connection to DB
func (c *Config) ReadFile() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
