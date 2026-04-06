package models

type User struct {
	ID        uint   `json:"id"`         // ID del usuario
	Name      string `json:"name"`       // Nombre del usuario
	Email     string `json:"email"`      // Correo electrónico del usuario
	Password  string `json:"password"`   // Contraseña del usuario (en un entorno real, esta debería ser hasheada)
	CreatedAt string `json:"created_at"` // Fecha de creación del usuario
	UpdatedAt string `json:"updated_at"` // Fecha de última actualización del usuario
}

type SignUpRequest struct {
	Name     string `json:"name" binding:"required"`     // Nombre del usuario (requerido)
	Email    string `json:"email" binding:"required"`    // Correo electrónico del usuario (requerido)
	Password string `json:"password" binding:"required"` // Contraseña del usuario (requerida)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`    // Correo electrónico del usuario (requerido)
	Password string `json:"password" binding:"required"` // Contraseña del usuario (requerida)
}
