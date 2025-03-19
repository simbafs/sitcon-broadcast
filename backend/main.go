package main

import (
	"embed"

	"backend/api"
	"backend/config"
	"backend/internal/fileserver"
	"backend/internal/logger"
	"backend/internal/staticfs"
	"backend/internal/token"
	"backend/models"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
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

	t := token.NewToken("token", "localhost:3000")

	r.POST("/verify", t.Verify())

	// TODO: replace DefaultConfig with my own config
	humaapi := humagin.New(r, huma.DefaultConfig("SITCON", "v1.0.0"))

	api.Route(humaapi, t)

	r.Use(t.ProtectRoute([]string{"/admin/card"}))
	fileserver.Route(r, static, Mode)

	log.Printf("Server is running at %s\n", c.Addr)
	return r.Run(c.Addr)
}

func main() {
	c := &config.Config{}
	c.SetDefault()
	c.FromEnv()

	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to init database: %v\n", err)
	}

	if err := run(c); err != nil {
		log.Printf("Oops, there's an error: %v\n", err)
	}
}
