package api

import (
	"backend/api/session"
	"backend/internal/logger"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api")

func Route(api huma.API) {
	session.Route(api)

	// TODO: 404
}
