package now

import (
	"io"
	"net/http"
	"strconv"

	"backend/middleware"
	"backend/models/now"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter, broadcast chan middleware.SSEMsg, t *middleware.TokenVerifyer, updateAll chan struct{}) {
	route := r.Group("/now")

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"now": now.GetNow(),
		})
	})

	route.POST("/", t.Auth, func(c *gin.Context) {
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
		updateAll <- struct{}{}
	})

	route.DELETE("/", t.Auth, func(c *gin.Context) {
		now.ClearNow()
		c.JSON(http.StatusOK, gin.H{"message": "cleared"})
		updateAll <- struct{}{}
	})
}
