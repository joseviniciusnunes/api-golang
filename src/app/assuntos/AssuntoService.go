package assuntos

import (
	"errors"
	"time"

	db "github.com/joseviniciusnunes/api-notificacao-golang/src/database"
	Response "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"gorm.io/gorm"
)

func CriarAssunto(dto *CriarAssuntoDtoRequest) Response.ResponseHttp {

	assuntoCreated := AssuntoModel{
		Titulo:       dto.Titulo,
		Conteudo:     dto.Conteudo,
		Icone:        dto.Icone,
		ChaveAcaoApp: dto.ChaveAcaoApp,
		Opcional:     dto.Opcional,
		Topico:       dto.Topico,
	}

	result := db.Con.Create(&assuntoCreated)

	if result.Error != nil {
		return Response.InternalServerError(result.Error.Error())
	}

	return Response.Created(assuntoCreated)
}

func ObterAssuntos() Response.ResponseHttp {
	result := []AssuntoModel{}
	returned := db.Con.Where("excluido_em IS NULL").Find(&result)
	if returned.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}
	return Response.Success(result)
}

func ObterAssunto(id int) Response.ResponseHttp {
	result := AssuntoModel{}
	returned := db.Con.First(&result, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum assunto foi encontrado")
	}
	return Response.Success(result)
}

func AlterarAssunto(dto *AlterarAssuntoDtoRequest, id int) Response.ResponseHttp {
	assunto := AssuntoModel{}
	returned := db.Con.First(&assunto, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum assunto foi encontrado")
	}

	assunto.Titulo = dto.Titulo
	assunto.Conteudo = dto.Conteudo
	assunto.Icone = dto.Icone
	assunto.ChaveAcaoApp = dto.ChaveAcaoApp
	assunto.Opcional = dto.Opcional
	assunto.Topico = dto.Topico

	returnedSave := db.Con.Save(&assunto)

	if returnedSave.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}

	return Response.Success(assunto)
}

func DeletarAssunto(id int) Response.ResponseHttp {
	assunto := AssuntoModel{}
	returned := db.Con.First(&assunto, id)
	if errors.Is(returned.Error, gorm.ErrRecordNotFound) {
		return Response.BadRequest("Nenhum assunto foi encontrado")
	}

	agora := time.Now()

	assunto.ExcluidoEm = &agora

	returnedSave := db.Con.Save(&assunto)

	if returnedSave.Error != nil {
		return Response.InternalServerError(returned.Error.Error())
	}

	return Response.SuccessMessage("Assunto deletado com sucesso")
}
