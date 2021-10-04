package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Con *gorm.DB

func ConnectToDatabase() error {
	mysqlString := os.Getenv("MYSQL_STRING")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	var err error
	Con, err = gorm.Open(mysql.Open(mysqlString), &gorm.Config{
		Logger: newLogger,
	})
	return err
}
