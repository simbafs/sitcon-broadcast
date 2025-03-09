package wrap

import (
	"net/http"

	"backend/internal/logger"

	"github.com/gin-gonic/gin"
)

// TODO: rename this package to a proper name

var log = logger.New("api")

func API(f func(c *gin.Context) any) gin.HandlerFunc {
	return func(c *gin.Context) {
		o := f(c)
		if len(c.Errors) > 0 {
			log.Println("err", c.Request.URL.Path, c.Errors.String())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": c.Errors.String(), // TODO: chagne to a machine readable format
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": o,
			})
		}
	}
}
