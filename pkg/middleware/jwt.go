package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Havens-blog/e-cas-service/internal/ecode"
	"github.com/Havens-blog/e-cas-service/pkg/jwt"
	"github.com/Havens-blog/e-cas-service/pkg/utils/response"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(auth) != 2 {
			response.Fail(c, ecode.TokenFailed, nil)
			c.Abort()
			return
		}
		token := auth[1]
		claims, err := jwt.Parse(token)
		if err != nil {
			response.Fail(c, ecode.TokenFailed, err)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
