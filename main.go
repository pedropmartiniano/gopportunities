package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropmartiniano/gopportunities/config"
	"github.com/pedropmartiniano/gopportunities/router"
)

var (
	logger *config.Logger = config.GetLogger()
)

func main() {
	r := gin.Default()

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Initialize(r)

	r.Run(":8081")
}
