package aplicativos

import (
	"strconv"

	res "github.com/joseviniciusnunes/api-notificacao-golang/src/helpers/ResponseHttp"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(router *echo.Echo) {
	router.GET("/api/v1/aplicativos", res.HandlerResponse(obterAplicativos))
	router.GET("/api/v1/aplicativos/:id", res.HandlerResponse(obterAplicativo))
	router.POST("/api/v1/aplicativos", res.HandlerResponse(criarAplicativo))
}

func obterAplicativos(response res.ResponseHttp, ctx echo.Context) res.ResponseHttp {
	return ObterAplicativos(response)
}

func obterAplicativo(response res.ResponseHttp, ctx echo.Context) res.ResponseHttp {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return response.InternalServerError(ctx.Param("id") + " é um ID inválido")
	}
	return ObterAplicativo(response, id)
}

func criarAplicativo(response res.ResponseHttp, ctx echo.Context) res.ResponseHttp {
	aplicativoDto := new(AplicativoDtoRequest)
	if errBind := ctx.Bind(aplicativoDto); errBind != nil {
		return response.InternalServerError(errBind.Error())
	}
	return CriarAplicativo(response, aplicativoDto)
}
