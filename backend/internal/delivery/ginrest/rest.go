package ginrest

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			var errorMessages []string
			for _, e := range c.Errors {
				errorMessages = append(errorMessages, e.Error())
			}

			log.Println("errors: ", errorMessages)
			log.Println(c.Writer.Written())

			if !c.Writer.Written() {
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": errorMessages,
				})
			}
		}
	}
}

func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "not found",
			})
		} else {
			c.String(http.StatusNotFound, "404 not found")
		}
	}
}
