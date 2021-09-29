package aplicativos

import (
	"errors"

	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	res "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"gorm.io/gorm"
)

func CriarAplicativo(response res.ResponseHttp, dto *AplicativoDtoRequest) res.ResponseHttp {

	aplicativoCreated := AplicativoModel{
		Nome:               dto.Nome,
		CredencialFirebase: dto.CredencialFirebase,
	}

	result := db.Con.Create(&aplicativoCreated)

	if result.Error != nil {
		return response.InternalServerError(result.Error.Error())
	}

	return response.Created(aplicativoCreated)
}

func ObterAplicativos(response res.ResponseHttp) res.ResponseHttp {
	result := []AplicativoModel{}
	returned := db.Con.Find(&result)
	if returned.Error != nil {
		return response.InternalServerError(returned.Error.Error())
	}
	return response.Success(result)
}

func ObterAplicativo(response res.ResponseHttp, id int) res.ResponseHttp {
	result := AplicativoModel{}
	returned := db.Con.First(&result, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return response.BadRequest("Nenhum aplicativo foi encontrado")
	}
	return response.Success(result)
}
