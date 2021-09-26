package aplicativos

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(router *echo.Echo) {
	router.GET("/api/v1/aplicativos", obterAplicativos)
	router.GET("/api/v1/aplicativos/{id}", obterAplicativo)
	router.POST("/api/v1/aplicativos", criarAplicativo)
}

func obterAplicativos(ctx echo.Context) error {
	result := ObterAplicativos()
	return ctx.JSON(http.StatusCreated, result)
	return nil
}

func obterAplicativo(ctx echo.Context) error {
	fmt.Println("obterAplicativo")
	return nil
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
