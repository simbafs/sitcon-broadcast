package token

import (
	"errors"
	"net/http"
	"strings"

	"backend/internal/logger"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
)

var log = logger.New("token")

type Token struct {
	token      string
	cookieName string
	domain     string
	maxAge     int
}

func NewToken(token string, domain string) *Token {
	return &Token{
		token:      token,
		cookieName: "token",
		domain:     domain,
		maxAge:     3600,
	}
}

// CheckToken is just a helper function to check if the sessio is valid
func (t *Token) CheckToken(c *gin.Context) bool {
	token, err := c.Cookie(t.cookieName)
	if err != nil {
		return false
	}

	return token == t.token
}

// RenewToken is a helper function to renew cookie
func (t *Token) RenewToken(c *gin.Context) {
	log.Printf("%#v", t)
	c.SetCookie(t.cookieName, t.token, t.maxAge, "/", t.domain, true, true)
}

type TokenBody struct {
	Token string `json:"token"`
}

// Verify is a gin api endpoint for /verify to check if the token in form is currect.
// It will set cookie if the token in post form is currect.
func (t *Token) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := TokenBody{}
		err := c.BindJSON(&res)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
			return
		}
		if res.Token != t.token {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error"})
			return
		}
		t.RenewToken(c)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

// AuthHuma add auth information to [huma.Operation]
func (t *Token) AuthHuma(api huma.API) func(op *huma.Operation) {
	return func(op *huma.Operation) {
		op.Middlewares = append(op.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
			token, err := huma.ReadCookie(ctx, t.cookieName)
			if err != nil {
				huma.WriteErr(api, ctx, http.StatusUnauthorized, "token miss", err)
				return
			}
			if err := token.Valid(); err != nil {
				huma.WriteErr(api, ctx, http.StatusUnauthorized, "token invalid", err)
				return
			}
			if token.Value != t.token {
				huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", errors.New("invalid token"))
				return
			}
			next(ctx)
		})
		op.Security = append(op.Security, map[string][]string{
			"token": {},
		})
	}
}

// ProtectRoute is a gin middleware to prevent unaresulthorized access to some static files
func (t *Token) ProtectRoute(routes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, route := range routes {
			if strings.HasPrefix(c.Request.URL.Path, route) {
				if !t.CheckToken(c) {
					log.Printf("Unauthorized access to %s from %s", c.Request.URL.Path, c.ClientIP())
					c.Redirect(http.StatusTemporaryRedirect, "/verify?redirect="+c.Request.URL.Path)
					return
				}
				t.RenewToken(c)
				c.Next()
				return
			}
		}
		c.Next()
	}
}
