package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativos"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/assuntos"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	"github.com/labstack/echo/v4"
)

func main() {

	godotenv.Load()

	errCon := database.ConnectToDatabase()
	if errCon != nil {
		fmt.Println(errCon.Error())
		panic("failed to connect database")
	}

	port := os.Getenv("PORT")

	router := echo.New()
	aplicativos.RegisterRoute(router)
	assuntos.RegisterRoute(router)
	fmt.Println("Listen on " + port)
	http.ListenAndServe(":"+port, router)
}
