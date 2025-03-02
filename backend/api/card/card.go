package card

import (
	_ "embed"
	"net/http"
	"time"

	"backend/logger"
	"backend/middleware"
	"backend/models/session"
	"backend/ticker"

	"github.com/gin-gonic/gin"
)

var log = logger.New("api/card")

type UpdateSession struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func Route(r gin.IRouter, t *middleware.TokenVerifyer, update chan ticker.Msg) {
	route := r.Group("/card")

	// get session
	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		if s, err := session.ReadByID(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	route.GET("/:room", func(c *gin.Context) {
		room := c.Param("room")

		if s, err := session.ReadCurrentByRoom(c.Request.Context(), room); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	// update session, the session must exist
	route.POST("/:id", t.VerifyToken, func(c *gin.Context) {
		id := c.Param("id")

		var u UpdateSession
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}

		if err := session.Update(c.Request.Context(), id, u.Start, u.End); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "filaed to update session"})
			return
		}

		update <- ticker.MsgCard
	})
}
