package repositories

import (
	"api-basico-dev/models"
	"context"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(
	ctx context.Context, // Se recibe el contexto para manejar la cancelación y los tiempos de espera
	user *models.User, // user SE RECIBE UN MODELO (se recibe como puntero para modificarlo directamente)
) error {

	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	res, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password)

	if err != nil {
		// Manejo del error (en un entorno real, se debería registrar el error y devolver un mensaje adecuado)
		return fmt.Errorf("Error al crear un nuevo usuario: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("Error al obtener el ID del nuevo usuario: %w", err)
	}

	user.ID = uint(id) // Se asigna el ID generado al modelo de usuario

	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = ?"
	res := r.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No se encontró el usuario
		}
		return nil, fmt.Errorf("Error al buscar el usuario por ID: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	res := r.db.QueryRowContext(ctx, query, email)

	var user models.User
	err := res.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, fmt.Errorf("Error al buscar el usuario por email: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	count := 0
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	res := r.db.QueryRowContext(ctx, query, email)

	err := res.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No se encontró el usuario
		}
		return false, fmt.Errorf("Error al buscar el usuario por email: %w", err)
	}

	return count > 0, nil
}
