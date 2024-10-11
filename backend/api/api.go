package api

import (
	"backend/api/card"
	"backend/api/now"
	"backend/api/room"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	api := r.Group("/api")
	// sse := middleware.NewSSE()

	card.Route(api)
	now.Route(api)
	room.Route(api)

	quit := make(chan struct{})
	go room.Ticker(quit)
}
