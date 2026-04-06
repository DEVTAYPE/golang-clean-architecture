package models

type Post struct {
	ID        uint   `json:"id"`         // ID del post
	UserId    uint   `json:"user_id"`    // ID del usuario que creó el post
	Title     string `json:"title"`      // Título del post
	Content   string `json:"content"`    // Contenido del post
	CreatedAt string `json:"created_at"` // Fecha de creación del post
	UpdatedAt string `json:"updated_at"` // Fecha de última actualización del post
}
