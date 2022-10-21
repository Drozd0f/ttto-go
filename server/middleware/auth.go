package middleware

import (
	"context"
	"net/http"

	"github.com/Drozd0f/ttto-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Context().Value("user") == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization token required",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func Auth(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("Authorization")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err == nil {
			ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "user", models.User{
				ID:       claims["ID"].(string),
				Username: claims["Username"].(string),
			}))
		}
		ctx.Next()
	}
}
