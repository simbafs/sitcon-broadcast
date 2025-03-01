package middleware

import (
	"net/http"

	"backend/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.New("middleware")

type TokenVerifyer struct {
	token  string
	domain string
}

func NewTokenVerifyer(token string, domain string) *TokenVerifyer {
	return &TokenVerifyer{
		token:  token,
		domain: domain,
	}
}

// Allow check if the token from cookie or query is valid. If yes and the cookie is from query, set the cookie.
func (t *TokenVerifyer) Allow(c *gin.Context) bool {
	token, err := c.Cookie("token")
	if err != nil || token != t.token {
		token := c.Query("token")
		if token == "" || token != t.token {
			return false
		}

		// expire in 7 days
		c.SetSameSite(http.SameSiteStrictMode)
		c.SetCookie("token", t.token, 7*24*60*60, "/", t.domain, false, true)
	}

	return true
}

// a gin middleware, if invalid it will redirect to /token
func (t *TokenVerifyer) ProtectRoute(routes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, route := range routes {
			if c.Request.URL.Path == route {
				if !t.Allow(c) {
					c.Redirect(http.StatusFound, "/token?redirect="+c.Request.URL.Path)
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

// A gin handler to provide a token verification endpoint
func (t *TokenVerifyer) VerifyToken(c *gin.Context) {
	if !t.Allow(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
	c.Abort()
}

// a gin middleware to let only valid token pass
func (t *TokenVerifyer) Auth(c *gin.Context) {
	if !t.Allow(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	c.Next()
}
