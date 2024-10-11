package now

import (
	"backend/middleware"
	"backend/models/now"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter, broadcast chan middleware.SSEMsg) {
	route := r.Group("/now")

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"now": now.GetNow(),
		})
	})

	route.POST("/", func(c *gin.Context) {
		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
		s := string(b)

		t, err := strconv.Atoi(s)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid time"})
			return
		}

		now.SetNow(t)
		c.JSON(http.StatusOK, gin.H{"time": t})
		broadcast <- middleware.SSEMsg{
			Name: "now",
			Data: now.GetNow(),
		}
	})

	route.DELETE("/", func(c *gin.Context) {
		now.ClearNow()
		c.JSON(http.StatusOK, gin.H{"message": "cleared"})
		broadcast <- middleware.SSEMsg{
			Name: "now",
			Data: now.GetNow(),
		}
	})
}
