package aplicativos

import (
	"errors"
	"time"

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
	returned := db.Con.Where("excluido_em IS NULL").Find(&result)
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

func AlterarAplicativo(dto *AlterarAplicativoDtoRequest, id int) Response.ResponseHttp {
	aplicativo := AplicativoModel{}
	returned := db.Con.First(&aplicativo, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum aplicativo foi encontrado")
	}

	aplicativo.Nome = dto.Nome
	aplicativo.CredencialFirebase = dto.CredencialFirebase

	returnedSave := db.Con.Save(&aplicativo)

	if returnedSave.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}

	return Response.Success(aplicativo)
}

func DeletarAplicativo(id int) Response.ResponseHttp {
	aplicativo := AplicativoModel{}
	returned := db.Con.First(&aplicativo, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum aplicativo foi encontrado")
	}

	agora := time.Now()

	aplicativo.ExcluidoEm = &agora

	returnedSave := db.Con.Save(&aplicativo)

	if returnedSave.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}

	return Response.SuccessMessage("Aplicativo deletado com sucesso")
}
