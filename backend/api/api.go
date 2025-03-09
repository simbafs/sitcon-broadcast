package api

import (
	"backend/api/session"
	"backend/internal/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.New("api")

func Route(router gin.IRouter) {
	api := router.Group("/api")

	session.Route(api)

	// TODO: 404
}
