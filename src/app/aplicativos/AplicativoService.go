package aplicativos

import (
	"errors"

	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	Response "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"gorm.io/gorm"
)

func CriarAplicativo(dto *CriarAplicativoDtoRequest) Response.ResponseHttp {

	aplicativoCreated := AplicativoModel{
		Nome:               dto.Nome,
		CredencialFirebase: dto.CredencialFirebase,
	}

	result := db.Con.Create(&aplicativoCreated)

	if result.Error != nil {
		return Response.InternalServerError(result.Error.Error())
	}

	return Response.Created(aplicativoCreated)
}

func ObterAplicativos() Response.ResponseHttp {
	result := []AplicativoModel{}
	returned := db.Con.Find(&result)
	if returned.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}
	return Response.Success(result)
}

func ObterAplicativo(id int) Response.ResponseHttp {
	result := AplicativoModel{}
	returned := db.Con.First(&result, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum aplicativo foi encontrado")
	}
	return Response.Success(result)
}
