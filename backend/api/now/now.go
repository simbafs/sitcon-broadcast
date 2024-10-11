package now

import (
	"backend/models/now"
	"backend/models/room"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter) {
	route := r.Group("/now")

	route.GET("/", func(c *gin.Context) {
		t := now.GetNow()
		c.JSON(http.StatusOK, gin.H{
			"now": t,
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
	})

	route.DELETE("/", func(c *gin.Context) {
		now.ClearNow()
		c.JSON(http.StatusOK, gin.H{"message": "cleared"})
	})

	route.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to parse room id",
			})
			return
		}

		if id >= len(room.Rooms) || id < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is out of range",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"now": room.Rooms[id].Time,
		})
	})
}
