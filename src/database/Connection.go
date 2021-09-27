package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Con *gorm.DB

func ConnectToDatabase() error {
	mysqlString := os.Getenv("MYSQL_STRING")
	var err error
	Con, err = gorm.Open(mysql.Open(mysqlString), &gorm.Config{})
	return err
}
