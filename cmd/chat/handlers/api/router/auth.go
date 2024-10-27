package router

import "github.com/gin-gonic/gin"

func SetAuthRouter(r *gin.Engine) {
	authV1 := r.Group("/v1/auth")
	{
		authV1.POST("/signup", func(ctx *gin.Context) {})
		authV1.POST("/login", func(ctx *gin.Context) {})
	}
}
