package handlers

import (
	"api-basico-dev/models"
	"api-basico-dev/server"
	"api-basico-dev/services"
	"net/http"
)

type NewUserHandler struct {
	userService *services.UserService //TODO: usar una interfaz en lugar de una estructura concreta para el servicio de usuarios
}

func NewNewUserHandler(userService *services.UserService) *NewUserHandler {
	return &NewUserHandler{userService: userService}
}

func (h *NewUserHandler) SignUpHandler(ctx *server.Context) {
	var req models.SignUpRequest // modelo para recibir los datos de la solicitud de registro

	// bindear los datos de la solicitud al modelo
	if err := ctx.BindJSON(&req); err != nil {
		RespondError(ctx, NewAppError("Datos de solicitud inválidos", 400))
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		RespondError(ctx, NewAppError("Todos los campos son obligatorios", 400))
		return
	}

	user, err := h.userService.SignUp(ctx.Ctx, req.Name, req.Email, req.Password)

	if err != nil {
		RespondError(ctx, NewAppError(err.Error(), 400))
		return
	}

	// enviar la respuesta al cliente con el usuario creado (sin la contraseña)
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Usuario creado exitosamente",
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
