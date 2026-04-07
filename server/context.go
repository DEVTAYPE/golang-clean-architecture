package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	RWriter http.ResponseWriter // para escribir la respuesta al cliente
	Request *http.Request       // para acceder a la información de la solicitud del cliente
	Ctx     context.Context     // para manejar el contexto de la solicitud, como cancelación o tiempo de espera
	UserUID uint                // para almacenar el UID del usuario autenticado, si es necesario
}

// Metodo para enviar una respuesta de texto al cliente
func (c *Context) Send(text string) {
	c.RWriter.Write([]byte(text))
}

// Metodo para enviar un codigo de estado HTTP al cliente
func (c *Context) SendStatus(statusCode int) {
	c.RWriter.WriteHeader(statusCode)
}

func (c *Context) JSON(
	statusCode int,
	data any,
) error {
	c.RWriter.Header().Set("Content-Type", "application/json")
	c.RWriter.WriteHeader(statusCode)

	err := json.NewEncoder(c.RWriter).Encode(data)

	if err != nil {
		http.Error(c.RWriter, "Error al tratar la respuesta", http.StatusInternalServerError)
		return err
	}

	return nil
}

func (c *Context) SetUserUID(uid uint) {
	c.UserUID = uid
}

func (c *Context) BindJSON(dest any) error {
	err := json.NewDecoder(c.Request.Body).Decode(dest)

	if err != nil {
		return err
	}

	return nil
}
