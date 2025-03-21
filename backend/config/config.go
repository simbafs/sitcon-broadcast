package config

import "os"

type Config struct {
	// for gin
	Addr string
	// for cookie
	Token string
	// for cookie
	Domain string
	// for ent
	DB string
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

	if c.DB == "" {
		c.DB = "./data/data.sqlite"
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

	if db, ok := os.LookupEnv("DB"); ok {
		c.DB = db
	}
}
