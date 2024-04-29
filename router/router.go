package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(router *gin.Engine) {
	initializeRoutes(router)
}
