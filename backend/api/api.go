package api

import (
	"backend/api/event"
	"backend/api/now"
	"backend/api/session"
	"backend/internal/logger"
	"backend/internal/token"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api")

func Route(api huma.API, t *token.Token) {
	api.OpenAPI().Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"token": {
			Type: "apiKey",
			In:   "cookie",
			Name: "token",
		},
		// NOTE: 不知道為什麼，OpenAPI 的 docs 界面雖然有個欄位讓我填 token，但是無論是按下「send api request」還是生成的 curl 命令，都沒有帶上 cookie
	}
	session.Route(huma.NewGroup(api, "/api/session"), t)
	now.Route(huma.NewGroup(api, "/api/now"), t)
	event.Route(huma.NewGroup(api, "/api/event"), t)

	// TODO: 404
}
