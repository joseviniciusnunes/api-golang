package main

import (
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativos"
	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
)

func main() {
	db.ConnectToDatabase()
	err := db.Con.AutoMigrate(&aplicativos.AplicativoModel{})
	if err != nil {
		panic(err.Error())
	}
}
