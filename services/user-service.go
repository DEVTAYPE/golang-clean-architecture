package services

import (
	"api-basico-dev/config"
	"api-basico-dev/helpers"
	"api-basico-dev/models"
	"api-basico-dev/repositories"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: Crear una interfaz para el servicio de usuarios, con métodos como CreateUser, GetUserByID, UpdateUser, DeleteUser, etc.

type UserService struct {
	// creamos un campo para el repositorio de usuarios, apuntando a la estructura UserRepository
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) generateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"expires_at": jwt.TimeFunc().Add(24 * time.Hour).Unix(), // el token expirará en 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *UserService) SignUp(
	ctx context.Context,
	name, email, password string,
) (*models.User, error) {

	if err := helpers.ValidateEmail(email); err != nil {
		return nil, err
	}

	if err := helpers.ValidatePassword(password); err != nil {
		return nil, err
	}

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

func (s *UserService) Login(
	ctx context.Context,
	email, password string,
) (string, error) {
	user, err := s.repo.FindByEmail(ctx, strings.ToLower(email))

	if err != nil {
		return "", fmt.Errorf("credenciales inválidas")
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return "", fmt.Errorf("credenciales inválidas")
	}

	token, err := s.generateToken(user.ID)

	if err != nil {
		return "", fmt.Errorf("error al generar el token: %v", err)
	}

	return token, nil
}
