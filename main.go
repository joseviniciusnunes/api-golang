package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	godotenv.Load()

	startDatabase()

	port := os.Getenv("PORT")

	router := mux.NewRouter()
	aplicativo.RegisterRoute(router)
	fmt.Println("Listen on " + port)
	http.ListenAndServe(":"+port, router)
}

func startDatabase() {
	var err error
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
}
