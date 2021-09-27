package aplicativos

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(router *echo.Echo) {
	router.GET("/api/v1/aplicativos", obterAplicativos)
	router.GET("/api/v1/aplicativos/:id", obterAplicativo)
	router.POST("/api/v1/aplicativos", criarAplicativo)
}

func obterAplicativos(ctx echo.Context) error {
	result, err := ObterAplicativos()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Erro ao buscar lista de aplicativos")
	}
	return ctx.JSON(http.StatusCreated, result)
}

func obterAplicativo(ctx echo.Context) error {
	var id int
	var err error
	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		return ctx.JSON(http.StatusBadRequest, ctx.Param("id")+" é um ID inválido")
	}
	result, errFind := ObterAplicativo(id)
	if errFind != nil {
		return ctx.JSON(http.StatusBadRequest, "Erro ao buscar aplicativo")
	}
	return ctx.JSON(http.StatusCreated, result)
}

func criarAplicativo(ctx echo.Context) error {
	aplicativoDto := new(AplicativoDtoRequest)
	if errBind := ctx.Bind(aplicativoDto); errBind != nil {
		return errBind
	}
	aplicativoCreated, errService := CriarAplicativo(aplicativoDto)
	if errService != nil {
		return errService
	}
	return ctx.JSON(http.StatusCreated, aplicativoCreated)
}
