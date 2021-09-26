package aplicativo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/aplicativos", obterAplicativos).Methods("GET")
	router.HandleFunc("/api/v1/aplicativos/{id}", obterAplicativo).Methods("GET")
}

func obterAplicativos(res http.ResponseWriter, req *http.Request) {
	fmt.Println("obterAplicativos")
}

func obterAplicativo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("obterAplicativo")
}
