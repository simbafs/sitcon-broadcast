package api

import (
	"backend/api/card"
	"backend/api/now"
	"backend/api/room"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	api := r.Group("/api")

	card.Route(api)
	now.Route(api)
	room.Route(api)

	api.GET("/sse", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		for i := 0; i < 5; i++ {
			data := fmt.Sprintf("data: {\"name\": \"test\", \"data\": %d}\n\n", i)
			_, err := c.Writer.WriteString(data)
			if err != nil {
				c.Error(err)
				return
			}

			c.Writer.Flush()

			time.Sleep(1 * time.Second)
		}
	})

	quit := make(chan struct{})
	go room.Ticker(quit)
}
