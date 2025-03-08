package api

import (
	"backend/api/card"
	"backend/api/countdown"
	"backend/api/now"
	"backend/api/special"
	"backend/middleware"
	"backend/ticker"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine, t *middleware.TokenVerifyer) {
	api := r.Group("/api")
	sse := middleware.NewSSE()

	update := make(chan ticker.Msg)

	card.Route(api, t, update)
	now.Route(api, t, update)
	countdown.Route(api, t, update)
	special.Route(api, t, update)

	api.GET("/sse", sse.GinHandler())

	quit := make(chan struct{})
	go ticker.Listen(sse.Message, quit, update)
}
