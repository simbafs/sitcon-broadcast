package special

import (
	"net/http"

	"backend/middleware"
	"backend/models/special"
	"backend/ticker"

	"github.com/gin-gonic/gin"
)

func Route(r gin.IRouter, t *middleware.TokenVerifyer, update chan ticker.Msg) {
	route := r.Group("/special")

	route.GET("/", func(c *gin.Context) {
		s, err := special.ReadAll(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, s)
	})

	route.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		s, err := special.Read(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Data(http.StatusOK, "application/json", []byte(s.Data))
	})

	route.PUT("/:id", t.Auth, func(c *gin.Context) {
		id := c.Param("id")
		var data string

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		s, err := special.Update(c, id, data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, s)
		update <- ticker.MsgSpecial(id)
	})
}
