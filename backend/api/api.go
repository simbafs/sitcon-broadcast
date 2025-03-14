package api

import (
	"backend/api/now"
	"backend/api/session"
	"backend/internal/logger"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api")

func Route(api huma.API) {
	session.Route(huma.NewGroup(api, "/api/session"))
	now.Route(huma.NewGroup(api, "/api/now"))

	// TODO: 404
}
