package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-restapi-jwt/utils/token"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
	}
}
