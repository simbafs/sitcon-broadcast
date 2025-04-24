package main

import (
	"embed"
	"time"

	"backend/api"
	"backend/internal/config"
	"backend/internal/logger"
	"backend/internal/token"
	"backend/models"
	"backend/models/now"
	"backend/sse"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-gonic/gin"
	"github.com/simbafs/kama"
)

// go embed ignore files begin with '_' or '.', 'all:' tells go embed to embed all files

//go:embed all:static/*
var rawStatic embed.FS

var (
	Mode       = "debug"
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

var log = logger.New("main")

func humaConfig() huma.Config {
	config := huma.DefaultConfig("SITCON Broadcast", "v1.0.0")

	config.Info.Description = "SITCON 製播組工具，目前包含直播字卡。"
	config.Info.Contact = &huma.Contact{
		Name: "SITCON 製播組 - 長條貓",
		URL:  "https://github.com/simbafs/sitcon-broadcast",
	}
	config.Info.License = &huma.License{
		Name:       "MIT",
		Identifier: "MIT",
		URL:        "https://github.com/simbafs/sitcon-broadcast/blob/main/LICENSE",
	}

	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"token": {
			Type: "apiKey",
			In:   "cookie",
			Name: "token",
		},
		// NOTE: 不知道為什麼，OpenAPI 的 docs 界面雖然有個欄位讓我填 token，但是無論是按下「send api request」還是生成的 curl 命令，都沒有帶上 cookie
	}

	// t, _ := json.MarshalIndent(config, "", "  ")
	// log.Printf("Huma config: %s\n", t)

	return config
}

func withSSE(r *gin.Engine) sse.Send {
	s := sse.New()
	r.GET("/sse", s.GinHandler())
	s.AddJob(time.NewTicker(time.Second), func() sse.Msg {
		return sse.Msg{
			Topic: []string{"now"},
			Data:  now.GetNow(),
		}
	})

	return s.Send
}

func withToken(r *gin.Engine, cfg *config.Config) *token.Token {
	t := token.NewToken(cfg.Token, cfg.Domain)
	r.POST("/verify", t.Verify())
	r.Use(t.ProtectRoute([]string{"/debug", "/event", "/counter/admin", "/card/admin"}))
	return t
}

func withKama(r *gin.Engine) {
	k := kama.New(rawStatic)
	r.Use(k.Gin())
}

func run(cfg *config.Config) error {
	gin.SetMode(Mode)
	r := gin.Default()

	send := withSSE(r)
	token := withToken(r, cfg)
	event, session, err := models.InitDB(cfg.DB)
	if err != nil {
		return err
	}
	api.Route(humagin.New(r, humaConfig()), api.NewHandler(send, token, event, session))

	withKama(r) // This must be put at the end of all middlewares

	log.Printf("Server is running at %s\n", cfg.Addr)
	return r.Run(cfg.Addr)
}

func main() {
	c := &config.Config{}
	c.SetDefault()
	c.FromEnv()

	if err := run(c); err != nil {
		log.Printf("Oops, there's an error: %v\n", err)
	}
}
