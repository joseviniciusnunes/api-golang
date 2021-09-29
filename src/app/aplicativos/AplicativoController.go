package aplicativos

import (
	"strconv"

	Response "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(router *echo.Echo) {
	router.GET("/api/v1/aplicativos", Response.HandlerResponse(obterAplicativos))
	router.GET("/api/v1/aplicativos/:id", Response.HandlerResponse(obterAplicativo))
	router.POST("/api/v1/aplicativos", Response.HandlerResponse(criarAplicativo))
}

func obterAplicativos(ctx echo.Context) Response.ResponseHttp {
	return ObterAplicativos()
}

func obterAplicativo(ctx echo.Context) Response.ResponseHttp {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return Response.InternalServerError(ctx.Param("id") + " é um ID inválido")
	}
	return ObterAplicativo(id)
}

func criarAplicativo(ctx echo.Context) Response.ResponseHttp {
	aplicativoDto := new(CriarAplicativoDtoRequest)
	if errBind := ctx.Bind(aplicativoDto); errBind != nil {
		return Response.InternalServerError(errBind.Error())
	}
	return CriarAplicativo(aplicativoDto)
}
