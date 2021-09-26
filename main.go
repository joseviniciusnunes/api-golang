package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativos"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	"github.com/labstack/echo/v4"
)

func main() {

	godotenv.Load()

	database.ConnectToDatabase()

	port := os.Getenv("PORT")

	router := echo.New()
	aplicativos.RegisterRoute(router)
	fmt.Println("Listen on " + port)
	http.ListenAndServe(":"+port, router)
}
