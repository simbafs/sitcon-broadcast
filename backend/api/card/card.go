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

	// get all sessions
	route.GET("/", func(c *gin.Context) {
		s, err := session.ReadAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, s)
	})

	// get session by id
	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")

		if s, err := session.ReadByID(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	// get current session in a room
	route.GET("/current/:room", func(c *gin.Context) {
		room := c.Param("room")

		if s, err := session.ReadCurrentByRoom(c.Request.Context(), room); err != nil {
			log.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	// update session, the session must exist
	route.PUT("/:room/:id", t.Auth, func(c *gin.Context) {
		room := c.Param("room")
		id := c.Param("id")

		var u UpdateSession
		if err := c.BindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := session.Update(c.Request.Context(), room, id, u.Start, u.End); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "updated"})

		update <- ticker.MsgCard{
			Room: room,
			ID:   id,
		}
	})
}
