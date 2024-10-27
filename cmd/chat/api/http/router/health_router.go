package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetHealthRouter(r *gin.Engine) {
	r.GET("/api/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
