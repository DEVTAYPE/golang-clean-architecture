package services

import (
	"api-basico-dev/helpers"
	"api-basico-dev/models"
	"api-basico-dev/repositories"
	"context"
	"fmt"
)

// TODO: Crear una interfaz para el servicio de usuarios, con métodos como CreateUser, GetUserByID, UpdateUser, DeleteUser, etc.

type UserService struct {
	// creamos un campo para el repositorio de usuarios, apuntando a la estructura UserRepository
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignUp(
	ctx context.Context,
	name, email, password string,
) (*models.User, error) {
	// verificar si el usuario ya existe en la base de datos
	exists, err := s.repo.EmailExists(ctx, email)

	if err != nil {
		return nil, err
	}

	// si el usuario ya existe, devolver un error indicando que el correo ya está registrado
	if exists {
		return nil, fmt.Errorf("el usuario con el correo %s ya existe", email)
	}

	passwordHash := helpers.HashPassword(password)
	password = passwordHash

	// crear un nuevo usuario utilizando el repositorio
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password, // el repositorio se encargará de hashear la contraseña
	}

	err = s.repo.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
