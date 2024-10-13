package api

import (
	"backend/api/card"
	"backend/api/countdown"
	"backend/api/now"
	"backend/middleware"
	"backend/models/session"
	"backend/ticker"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	api := r.Group("/api")
	sse := middleware.NewSSE()

	api.GET("/session", func(c *gin.Context) {
		c.JSON(http.StatusOK, session.Data.Rooms)
	})

	card.Route(api, sse.Message)
	now.Route(api, sse.Message)
	countdown.Route(api, sse.Message)

	api.GET("/sse", sse.GinHandler())

	quit := make(chan struct{})
	go ticker.Listen(sse.Message, quit)
}
