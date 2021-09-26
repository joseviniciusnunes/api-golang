package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joseviniciusnunes/api-notificacao-golang/src/app/aplicativo"
)

func main() {
	router := mux.NewRouter()
	aplicativo.RegisterRoute(router)
	http.ListenAndServe(":3000", router)
}
