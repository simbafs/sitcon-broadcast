package card

import (
	"backend/middleware"
	"backend/models/session"
	_ "embed"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter, broadcast chan middleware.SSEMsg, t *middleware.TokenVerifyer) {
	route := r.Group("/card")

	// get session
	route.GET("/:room/:idx", func(c *gin.Context) {
		room := c.Param("room")
		idxStr := c.Param("idx")

		idx, err := strconv.Atoi(idxStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
			return
		}

		if s, ok := session.Data.Rooms.Get(room, idx); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	route.GET("/:room", func(c *gin.Context) {
		room := c.Param("room")

		if r, ok := session.Data.Rooms[room]; !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		} else if r, ok := r.GetNow(); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "no session now"})
			return
		} else {
			c.JSON(http.StatusOK, r)
		}
	})

	// update session, the session must exist
	route.POST("/:room/:idx", t.VerifyToken, func(c *gin.Context) {
		room := c.Param("room")
		idxStr := c.Param("idx")

		idx, err := strconv.Atoi(idxStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid index"})
			return
		}

		var s session.SessionItem
		if err := c.BindJSON(&s); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
			return
		}

		if _, ok := session.Data.Rooms.Get(room, idx); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}

		if session.Data.Rooms[room].IsOverlap(idx, s.Start, s.End) {
			c.JSON(http.StatusConflict, gin.H{"error": "room is not free at this time"})
			return
		}

		session.Data.Rooms[room][idx] = s

		c.JSON(http.StatusOK, s)
		broadcast <- middleware.SSEMsg{
			Name: "card-" + room,
			Data: s,
		}
	})

	// create a session, the time must not overlap with others
	route.PUT("/:room", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
		// room := c.Param("room")
		//
		// r, ok := data.Sessionss[room]
		// if !ok {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		// 	return
		// }
		//
		// var s session.SessionsItem
		// if err := c.BindJSON(&s); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		// 	return
		// }
		//
		// if r.IsOverlap(s.Start, s.End) {
		// 	c.JSON(http.StatusConflict, gin.H{"error": "room is not free at this time"})
		// 	return
		// }
		//
		// s.ID = data.GetNextID()
		//
		// data.Sessionss[room] = append(data.Sessions[room], s)
		//
		// slices.SortFunc(data.Sessionss[room], func(a, b session.SessionItem) int {
		// 	return int(a.Start - b.Start)
		// })
		//
		// c.JSON(http.StatusOK, s)
	})
}
