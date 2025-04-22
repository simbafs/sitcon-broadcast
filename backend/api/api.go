package api

import (
	"backend/api/counter"
	"backend/api/event"
	"backend/api/now"
	"backend/api/session"
	"backend/internal/logger"
	"backend/internal/token"
	"backend/sse"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api")

func Route(api huma.API, t *token.Token, send chan sse.Msg) {
	session.Route(huma.NewGroup(api, "/api/session"), t, send)
	now.Route(huma.NewGroup(api, "/api/now"), t, send)
	event.Route(huma.NewGroup(api, "/api/event"), t)
	counter.Route(huma.NewGroup(api, "/api/counter"), t, send)
}
