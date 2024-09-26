package card

import (
	"backend/models/now"
	"backend/models/session"
	"bytes"
	_ "embed"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//go:embed sessions.json
var file []byte

func Route(r gin.IRouter) {
	data, err := session.GetSessions(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}

	r.GET("/session", func(c *gin.Context) {
		c.JSON(http.StatusOK, data.Sessions)
	})

	r.GET("/idMap", func(c *gin.Context) {
		c.JSON(http.StatusOK, data.IDMap)
	})

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

		if s, ok := data.Sessions.Get(room, idx); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		} else {
			c.JSON(http.StatusOK, s)
		}
	})

	route.GET("/:room", func(c *gin.Context) {
		room := c.Param("room")

		if _, ok := data.Sessions[room]; !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}

		nowTime := now.GetNow()
		nowSession := data.Sessions[room][0]

		for _, s := range data.Sessions[room] {
			if s.Start > nowTime {
				break
			}

			nowSession = s
		}

		if nowSession.ID == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		}

		c.JSON(http.StatusOK, nowSession)
	})

	// update session, the session must exist
	route.POST("/:room/:idx", func(c *gin.Context) {
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

		if _, ok := data.Sessions.Get(room, idx); !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
			return
		}

		if data.Sessions[room].IsOverlap(s.Start, s.End) {
			c.JSON(http.StatusConflict, gin.H{"error": "room is not free at this time"})
			return
		}

		data.Sessions[room][idx] = s

		c.JSON(http.StatusOK, s)
	})

	// create a session, the time must not overlap with others
	route.PUT("/:room", func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
		// room := c.Param("room")
		//
		// r, ok := data.Sessions[room]
		// if !ok {
		// 	c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		// 	return
		// }
		//
		// var s session.SessionItem
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
		// data.Sessions[room] = append(data.Sessions[room], s)
		//
		// slices.SortFunc(data.Sessions[room], func(a, b session.SessionItem) int {
		// 	return int(a.Start - b.Start)
		// })
		//
		// c.JSON(http.StatusOK, s)
		// TODO: rebuild idMap
	})
}
