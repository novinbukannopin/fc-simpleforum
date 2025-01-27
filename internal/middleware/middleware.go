package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/novinbukannopin/fc-simple-forum/internal/configs"
	"github.com/novinbukannopin/fc-simple-forum/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			_ = c.AbortWithError(http.StatusUnauthorized, errors.New("missing Authorization header"))
			return
		}

		userId, username, err := jwt.ValidatedToken(header, secretKey)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
