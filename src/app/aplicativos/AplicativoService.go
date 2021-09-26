package aplicativos

import (
	"fmt"

	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
)

func CriarAplicativo(dto *AplicativoDtoRequest) (AplicativoModel, error) {

	aplicativoCreated := AplicativoModel{
		Nome:               dto.Nome,
		CredencialFirebase: dto.CredencialFirebase,
	}

	result := db.Con.Create(&aplicativoCreated)

	fmt.Println(aplicativoCreated.Id)
	fmt.Println(result.Error)

	return aplicativoCreated, nil
}

func ObterAplicativos() []AplicativoModel {
	result := []AplicativoModel{}
	db.Con.Find(&result)
	return result
}
