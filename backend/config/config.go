package config

import "os"

type Config struct {
	Addr   string
	Token  string
	Domain string
}

func (c *Config) SetDefault() {
	if c.Addr == "" {
		c.Addr = ":3000"
	}

	if c.Token == "" {
		c.Token = "token"
	}

	if c.Domain == "" {
		c.Domain = "localhost"
	}
}

func (c *Config) FromEnv() {
	if addr, ok := os.LookupEnv("ADDR"); ok {
		c.Addr = addr
	}

	if token, ok := os.LookupEnv("TOKEN"); ok {
		c.Token = token
	}

	if domain, ok := os.LookupEnv("DOMAIN"); ok {
		c.Domain = domain
	}
}
