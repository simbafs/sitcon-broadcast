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
	session.Route(huma.NewGroup(api, "/api/session"), t)
	now.Route(huma.NewGroup(api, "/api/now"), t)
	event.Route(huma.NewGroup(api, "/api/event"), t)
	// TODO: 404
}
