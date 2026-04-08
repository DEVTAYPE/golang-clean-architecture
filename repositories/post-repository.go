package repositories

import (
	"api-basico-dev/models"
	"context"
	"database/sql"
	"fmt"
)

// Se crea un struct para el repositorio de Post, que tendrá una conexión a la base de datos
type PostRepository struct {
	db *sql.DB // Conexión a la base de datos
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(
	ctx context.Context, // Se recibe el contexto para manejar la cancelación y los tiempos de espera
	post *models.Post, // post SE RECIBE UN MODELO (se recibe como puntero para modificarlo directamente)
) error {

	query := "INSERT INTO posts (title, content, user_id) VALUES (?, ?, ?)"
	res, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.UserId)

	if err != nil {
		return fmt.Errorf("Error al crear un nuevo post: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el ID del nuevo post: %w", err)
	}

	post.ID = uint(id) // Se asigna el ID generado al modelo de post

	return nil
}

func (r *PostRepository) FindByID(ctx context.Context, id uint) (*models.Post, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE id = ?"
	res := r.db.QueryRowContext(ctx, query, id)

	var post models.Post
	err := res.Scan(&post.ID, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró el post
		}
		return nil, fmt.Errorf("Error al buscar el post por ID: %w", err)
	}

	return &post, nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]*models.Post, error) {
	query := "SELECT id, title, content, user_id, created_at, updated_at FROM posts"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Error al buscar los posts: %w", err)
	}

	defer rows.Close() // Se asegura de cerrar las filas después de usarlas para liberar recursos

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Error al escanear el post: %w", err)
		}
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error al iterar sobre los posts: %w", err)
	}

	return posts, nil
}

func (r *PostRepository) FindByUserID(ctx context.Context, userId uint) ([]*models.Post, error) {

	query := "SELECT id, title, content, user_id, created_at, updated_at FROM posts WHERE user_id = ?"

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("Error al buscar los posts por ID de usuario: %w", err)
	}

	defer rows.Close()

	var posts []*models.Post
	// rows.Next() se utiliza para iterar sobre los resultados de la consulta. En cada iteración, se escanea el resultado en un nuevo modelo de Post y se agrega a la lista de posts.
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserId, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Error al escanear el post: %w", err)
		}
		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Error al iterar sobre los posts: %w", err)
	}

	return posts, nil
}

func (r *PostRepository) Update(ctx context.Context, post *models.Post) error {
	query := "UPDATE posts SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)

	if err != nil {
		return fmt.Errorf("Error al actualizar el post: %w", err)
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, postId uint) error {
	query := "DELETE FROM posts WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, postId)

	if err != nil {
		return fmt.Errorf("Error al eliminar el post: %w", err)
	}

	return nil
}
