package main

import (
	"embed"
	"log"

	"backend/api"
	"backend/config"
	"backend/internal/fileserver"
	"backend/internal/staticfs"
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

var logger = log.New(gin.DefaultWriter, "[main] ", log.LstdFlags|log.Lmsgprefix)

func run(c *config.Config) error {
	gin.SetMode(Mode)
	r := gin.Default()

	t := middleware.NewTokenVerifyer(c.Token, c.Domain)

	r.Use(t.ProtectRoute([]string{"/newCard/admin", "/countdown/admin", "/card/admin", "/debug"}))
	r.GET("/verify", t.VerifyToken)

	api.Route(r, t)
	fileserver.Route(r, static, Mode)

	logger.Printf("Server is running at %s\n", c.Addr)
	return r.Run(c.Addr)
}

func main() {
	c := &config.Config{}
	c.SetDefault()
	c.FromEnv()

	if err := run(c); err != nil {
		logger.Printf("Oops, there's an error: %v\n", err)
	}
}
