package now

import (
	"net/http"

	"backend/middleware"
	"backend/models/now"
	"backend/ticker"

	"github.com/gin-gonic/gin"
)

type NowBody struct {
	Now now.Now `json:"now"`
}

func Route(r gin.IRouter, t *middleware.TokenVerifyer, update chan ticker.Msg) {
	route := r.Group("/now")

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, NowBody{
			Now: now.Read(),
		})
	})

	route.PUT("/", t.Auth, func(c *gin.Context) {
		t := NowBody{}

		if err := c.BindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		now.Update(t.Now)

		c.JSON(http.StatusOK, gin.H{"message": "updated"})
		update <- ticker.MsgNow{}
	})

	route.DELETE("/", t.Auth, func(c *gin.Context) {
		now.Delete()
		c.JSON(http.StatusOK, gin.H{"message": "cleared"})
		update <- ticker.MsgNow{}
	})
}
