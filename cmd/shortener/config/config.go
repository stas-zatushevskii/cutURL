package config

import (
	"flag"
	"os"
	"strings"
)

type Config struct {
	ServerURL string
	BaseURL   string
}

func (c *Config) ParseFlags() {
	flag.StringVar(&c.ServerURL, "a", "127.0.0.1:8080", "address for server")
	flag.StringVar(&c.BaseURL, "b", "http://127.0.0.1:8080/", "address for base url of shortener")
	flag.Parse()
}

func (c *Config) ParseEnv() {
	if envServerURL := os.Getenv("SERVER_ADDRESS"); envServerURL != "" {
		c.ServerURL = envServerURL
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		c.BaseURL = envBaseURL
	}
}

func (c *Config) ValidateBaseURL() {
	if !strings.HasSuffix(c.BaseURL, "/") {
		c.BaseURL += "/"
	}
}

func LoadConfig() *Config {
	cfg := &Config{}
	cfg.ParseFlags()
	cfg.ParseEnv()
	cfg.ValidateBaseURL()
	return cfg
}
