package aplicativo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/aplicativos", obterAplicativos).Methods("GET")
	router.HandleFunc("/api/v1/aplicativos/{id}", obterAplicativo).Methods("GET")
	router.HandleFunc("/api/v1/aplicativos", criarAplicativo).Methods("POST")
}

func obterAplicativos(res http.ResponseWriter, req *http.Request) {
	fmt.Println("obterAplicativos")
}

func obterAplicativo(res http.ResponseWriter, req *http.Request) {
	fmt.Println("obterAplicativo")
}

func criarAplicativo(res http.ResponseWriter, req *http.Request) {
	var aplicativoDto AplicativoDtoRequest
	err := json.NewDecoder(req.Body).Decode(&aplicativoDto)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	errValid := aplicativoDto.IsValid()
	if errValid != nil {
		http.Error(res, errValid.Error(), http.StatusBadRequest)
		return
	}
	CriarAplicativo(aplicativoDto)
}
