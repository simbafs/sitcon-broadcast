package main

import (
	"embed"

	"backend/api"
	"backend/config"
	"backend/internal/fileserver"
	"backend/internal/staticfs"
	"backend/logger"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// go embed ignore files begin with '_' or '.', 'all:' tells go embed to embed all files

//go:embed all:static/*
var rawStatic embed.FS

var static = staticfs.NewStatic(rawStatic, "static")

var (
	Mode       = "debug"
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

var log = logger.New("main")

func run(c *config.Config) error {
	gin.SetMode(Mode)
	r := gin.Default()

	t := middleware.NewTokenVerifyer(c.Token, c.Domain)

	r.Use(t.ProtectRoute([]string{"/card/admin", "/countdown/admin", "/debug"}))
	r.GET("/verify", t.VerifyToken)

	api.Route(r, t)
	fileserver.Route(r, static, Mode)

	log.Printf("Server is running at %s\n", c.Addr)
	return r.Run(c.Addr)
}

func main() {
	c := &config.Config{}
	c.SetDefault()
	c.FromEnv()

	if err := run(c); err != nil {
		log.Printf("Oops, there's an error: %v\n", err)
	}
}
