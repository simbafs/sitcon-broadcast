package ginrest

import (
	"log"
	"net/http"
	"strings"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	event usecase.Event
}

func New(event usecase.Event) *Gin {
	return &Gin{
		event: event,
	}
}

func (g *Gin) Route(r gin.IRouter) {
	r.POST("/event", g.Create)
	r.DELETE("/event/:name", g.Delete)
	r.GET("/event", g.List)
	r.GET("/event/:name", g.Get)
	r.PUT("/event/:name", g.UpdateScript)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var errorMessages []string
			for _, e := range c.Errors {
				errorMessages = append(errorMessages, e.Error())
			}

			log.Println("errors: ", errorMessages)
			log.Println(c.Writer.Written())

			if !c.Writer.Written() {
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": errorMessages,
				})
			}
		}
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "not found",
			})
		} else {
			c.String(http.StatusNotFound, "404 not found")
		}
	}
}
