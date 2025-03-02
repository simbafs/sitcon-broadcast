package countdown

import (
	"log"
	"net/http"

	"backend/middleware"
	"backend/models/room"
	"backend/ticker"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter, t *middleware.TokenVerifyer, update chan ticker.Msg) {
	route := r.Group("/countdown")

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"rooms": room.Rooms,
		})
	})

	route.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")

		targetRoom, ok := room.Rooms[name]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "room not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"room": targetRoom,
		})
	})

	route.POST("/:name", t.Auth, func(c *gin.Context) {
		name := c.Param("name")
		if _, ok := room.Rooms[name]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "room not found",
			})
			return
		}

		r := room.Room{}

		if err := c.BindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		log.Printf("update room %s to %#v\n", name, r)

		room.Rooms[name] = r
		c.JSON(http.StatusOK, gin.H{
			"message": "success update room",
		})
		update <- ticker.MsgCountdown
	})
}
