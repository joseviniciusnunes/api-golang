package assuntos

import (
	"fmt"
	"strconv"

	Response "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(router *echo.Echo) {
	router.GET("/api/v1/assuntos", Response.HandlerResponse(obterAssuntos))
	router.GET("/api/v1/assuntos/:id", Response.HandlerResponse(obterAssunto))
	router.POST("/api/v1/assuntos", Response.HandlerResponse(criarAssunto))
	router.PUT("/api/v1/assuntos/:id", Response.HandlerResponse(alterarAssunto))
	router.DELETE("/api/v1/assuntos/:id", Response.HandlerResponse(deletarAssunto))
}

func obterAssuntos(ctx echo.Context) Response.ResponseHttp {
	return ObterAssuntos()
}

func obterAssunto(ctx echo.Context) Response.ResponseHttp {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return Response.InternalServerError(ctx.Param("id") + " é um ID inválido")
	}
	return ObterAssunto(id)
}

func criarAssunto(ctx echo.Context) Response.ResponseHttp {
	assuntoDto := new(CriarAssuntoDtoRequest)
	if errBind := ctx.Bind(assuntoDto); errBind != nil {
		return Response.InternalServerError(errBind.Error())
	}
	if errsValidate := assuntoDto.IsValid(); len(errsValidate) > 0 {
		fmt.Println(errsValidate)
		return Response.BadRequestList(errsValidate)
	}
	return CriarAssunto(assuntoDto)
}

func alterarAssunto(ctx echo.Context) Response.ResponseHttp {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return Response.InternalServerError(ctx.Param("id") + " é um ID inválido")
	}
	assuntoDto := new(AlterarAssuntoDtoRequest)
	if errBind := ctx.Bind(assuntoDto); errBind != nil {
		return Response.InternalServerError(errBind.Error())
	}
	return AlterarAssunto(assuntoDto, id)
}

func deletarAssunto(ctx echo.Context) Response.ResponseHttp {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return Response.InternalServerError(ctx.Param("id") + " é um ID inválido")
	}
	return DeletarAssunto(id)
}
