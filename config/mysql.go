package config

import (
	"github.com/pedropmartiniano/gopportunities/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger()

	const mysqlDSN string = "root:1234@tcp(127.0.0.1:3306)/gopportunities"

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: mysqlDSN}), &gorm.Config{})

	if err != nil {
		logger.Errorf("mysql opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
