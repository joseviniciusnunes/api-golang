package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativo"
)

func main() {

	godotenv.Load()

	port := os.Getenv("PORT")

	router := mux.NewRouter()
	aplicativo.RegisterRoute(router)
	fmt.Println("Listen on " + port)
	http.ListenAndServe(":"+port, router)
}
