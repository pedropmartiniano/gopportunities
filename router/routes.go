package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"msg": "GET Opening",
			})
		})
	}
}
