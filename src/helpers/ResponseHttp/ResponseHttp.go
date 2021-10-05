package ResponseHttp

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type ResponseHttp struct {
	Code    int
	Body    interface{}
	Error   error
	Errors  []error
	Headers map[string]string
}

type ErrorResponse struct {
	Status   int      `json:"status"`
	Message  string   `json:"message,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func Success(body interface{}) ResponseHttp {
	var res ResponseHttp
	res.Code = 200
	res.Body = body
	res.Error = nil
	return res
}

func SuccessMessage(message string) ResponseHttp {
	var res ResponseHttp
	res.Code = 200
	res.Body = MessageResponse{
		Message: message,
	}
	res.Error = nil
	return res
}

func Created(body interface{}) ResponseHttp {
	var res ResponseHttp
	res.Code = 201
	res.Body = body
	res.Error = nil
	return res
}

func BadRequest(message string) ResponseHttp {
	var res ResponseHttp
	res.Code = 400
	res.Error = errors.New(message)
	return res
}

func BadRequestList(messages []error) ResponseHttp {
	var res ResponseHttp
	res.Code = 400
	res.Errors = messages
	return res
}

func Unauthorized(message string) ResponseHttp {
	var res ResponseHttp
	res.Code = 401
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func Forbidden(message string) ResponseHttp {
	var res ResponseHttp
	res.Code = 403
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func InternalServerError(message string) ResponseHttp {
	var res ResponseHttp
	res.Code = 500
	res.Body = nil
	res.Error = errors.New(message)
	return res
}

func HandlerResponse(fn func(echo.Context) ResponseHttp) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		returned := fn(ctx)

		if returned.Code == 500 {
			error := ErrorResponse{
				Status:  returned.Code,
				Message: "Infelizmente estamos com problemas internos, já estamos trabalhando para resolve-los",
			}
			return ctx.JSON(returned.Code, error)
		}

		if returned.Code >= 400 {
			if len(returned.Errors) > 0 {
				var error ErrorResponse
				error.Status = returned.Code
				for _, e := range returned.Errors {
					error.Messages = append(error.Messages, e.Error())
				}
				return ctx.JSON(returned.Code, error)
			} else {
				var error ErrorResponse
				error.Status = returned.Code
				error.Message = returned.Error.Error()
				return ctx.JSON(returned.Code, error)
			}
		}

		return ctx.JSON(returned.Code, returned.Body)
	}
}
