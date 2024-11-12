package middlewares

import (
	"net/http"

	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[7:]

		isValid, claims := utils.IsTokenValid(token)

		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			ctx.Set("user", claims)

			ctx.Next()
		}

	}
}
