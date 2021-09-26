package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Con *gorm.DB

func ConnectToDatabase() {
	mysqlString := os.Getenv("MYSQL_STRING")
	var err error
	Con, err = gorm.Open(mysql.Open(mysqlString), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
}
