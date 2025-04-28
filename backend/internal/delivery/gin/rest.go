package gin

import (
	"backend/internal/delivery"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	event delivery.EventUsecase
}

func New(event delivery.EventUsecase) *Gin {
	return &Gin{
		event: event,
	}
}

func (g *Gin) Route(r *gin.RouterGroup) {
	r.POST("/event", g.Create)
	r.DELETE("/event/:name", g.Delete)
	r.GET("/event", g.List)
	r.GET("/event/:name", g.Get)
	r.PUT("/event/:name", g.UpdateScript)
}
