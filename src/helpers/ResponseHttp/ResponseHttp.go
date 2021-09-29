package ResponseHttp

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type ResponseHttp struct {
	Code  int
	Body  interface{}
	Error error
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (res ResponseHttp) Success(body interface{}) ResponseHttp {
	res.Code = 200
	res.Body = body
	res.Error = nil
	return res
}

func (res ResponseHttp) Created(body interface{}) ResponseHttp {
	res.Code = 201
	res.Body = body
	res.Error = nil
	return res
}

func (res ResponseHttp) BadRequest(message string) ResponseHttp {
	res.Code = 400
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func (res ResponseHttp) Unauthorized(message string) ResponseHttp {
	res.Code = 401
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func (res ResponseHttp) InternalServerError(message string) ResponseHttp {
	res.Code = 500
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func HandlerResponse(fn func(ResponseHttp, echo.Context) ResponseHttp) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var response ResponseHttp
		returned := fn(response, ctx)

		if returned.Code == 500 {
			error := ErrorResponse{
				Status:  returned.Code,
				Message: "Infelizmente estamos com problemas internos, já estamos trabalhando para resolve-lo",
			}
			return ctx.JSON(returned.Code, error)
		}

		if returned.Code >= 400 {
			error := ErrorResponse{
				Status:  returned.Code,
				Message: returned.Error.Error(),
			}
			return ctx.JSON(returned.Code, error)
		}

		return ctx.JSON(returned.Code, returned.Body)
	}
}
