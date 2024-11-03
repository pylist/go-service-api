package middleware

import (
	"go-service-api/pkg"
	"go-service-api/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.Request.Header.Get("Authorization")
		info := strings.Split(authorization, " ")
		if len(info) < 2 {
			response.ErrorCode(c, response.StatusErrTokenNotFound)
			c.Abort()
			return
		}
		token := info[1]
		claims, err := pkg.NewToken().ValidatingToken(token)
		if err != nil {
			response.ErrorCode(c, response.StatusErrTokenInvalid)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
