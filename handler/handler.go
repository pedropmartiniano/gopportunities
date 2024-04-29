package handler

import (
	"github.com/pedropmartiniano/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger = config.GetLogger()
	db     *gorm.DB
)

func InitializeHandler() {
	db = config.GetMySQL()
}
