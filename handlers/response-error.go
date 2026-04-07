package handlers

import (
	"api-basico-dev/server"
	"net/http"
)

type ResponseError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AppError struct {
	Message string
	Code    int
}

// nos sirve para convertir el error a string
func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(message string, code int) *AppError {
	return &AppError{
		Message: message,
		Code:    code,
	}
}

func RespondError(ctx *server.Context, appErr *AppError) {
	responseErr := ResponseError{
		Error:   http.StatusText(appErr.Code),
		Message: appErr.Message,
		Code:    appErr.Code,
	}

	// enviamos la respuesta de error al cliente en formato JSON
	ctx.JSON(appErr.Code, responseErr)
}
