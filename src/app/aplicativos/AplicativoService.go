package aplicativos

import (
	"errors"

	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	"gorm.io/gorm"
)

func CriarAplicativo(dto *AplicativoDtoRequest) (AplicativoModel, error) {

	aplicativoCreated := AplicativoModel{
		Nome:               dto.Nome,
		CredencialFirebase: dto.CredencialFirebase,
	}

	result := db.Con.Create(&aplicativoCreated)

	return aplicativoCreated, result.Error
}

func ObterAplicativos() ([]AplicativoModel, error) {
	result := []AplicativoModel{}
	response := db.Con.Find(&result)
	return result, response.Error
}

func ObterAplicativo(id int) (AplicativoModel, error) {
	result := AplicativoModel{}
	response := db.Con.First(&result, id)
	if errors.Is(response.Error, gorm.ErrRecordNotFound) {
		return result, errors.New("Nenhum aplicativo foi encontrado")
	}
	return result, response.Error
}
