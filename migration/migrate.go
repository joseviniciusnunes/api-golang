package main

import (
	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativos"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
)

func main() {
	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		panic(errEnv.Error())
	}

	errCon := database.ConnectToDatabase()
	if errCon != nil {
		panic(errCon.Error())
	}
	err := db.Con.AutoMigrate(&aplicativos.AplicativoModel{})
	if err != nil {
		panic(err.Error())
	}
}
