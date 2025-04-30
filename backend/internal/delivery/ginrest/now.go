package ginrest

import (
	"net/http"

	"backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Now struct {
	now usecase.Now
}

func NewNow(r gin.IRouter, now usecase.Now) *Now {
	n := &Now{
		now: now,
	}

	r.GET("/", n.Get)
	r.PUT("/", n.Set)
	r.DELETE("/", n.Reset)

	return n
}

func (n *Now) Get(c *gin.Context) {
	output := n.now.Get()
	c.JSON(http.StatusOK, output)
}

func (n *Now) Set(c *gin.Context) {
	var input usecase.NowInput
	err := c.BindJSON(&input)
	if err != nil {
		c.Error(err)
		return
	}

	output := n.now.Set(&input)
	c.JSON(http.StatusOK, output)
}

func (n *Now) Reset(c *gin.Context) {
	output := n.now.Reset()
	c.JSON(http.StatusOK, output)
}
