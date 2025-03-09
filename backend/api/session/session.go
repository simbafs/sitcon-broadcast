package session

import (
	"backend/api/wrap"
	"backend/models/session"

	"github.com/gin-gonic/gin"
)

func Route(router gin.IRouter) {
	api := router.Group("/session")

	// get current session in room
	api.GET("/:room", wrap.API(func(c *gin.Context) any {
		room := c.Param("room")

		s, err := session.GetCurrent(c, room)
		if err != nil {
			c.Error(err)
			return nil
		}

		return s
	}))

	// get a session by id in the room
	api.GET("/:room/:id", wrap.API(func(c *gin.Context) any {
		room := c.Param("room")
		id := c.Param("id")

		if id == "all" {
			s, err := session.GetAllInRoom(c, room)
			if err != nil {
				c.Error(err)
				return nil
			}
			return s
		} else {
			s, err := session.Get(c, room, id)
			if err != nil {
				c.Error(err)
				return nil
			}
			return s
		}
	}))
}
