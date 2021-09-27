package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativos"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
)

func main() {
	godotenv.Load("../.env")
	errCon := database.ConnectToDatabase()
	if errCon != nil {
		fmt.Println(errCon.Error())
		panic("failed to connect database")
	}
	err := db.Con.AutoMigrate(&aplicativos.AplicativoModel{})
	if err != nil {
		panic(err.Error())
	}
}
