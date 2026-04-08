package services

import (
	"api-basico-dev/models"
	"api-basico-dev/repositories"
	"context"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(
	repo *repositories.PostRepository,
) *PostService {
	return &PostService{repo: repo}
}

// Aquí puedes agregar métodos para manejar la lógica de negocio relacionada con los posts, como crear un post, obtener posts, actualizar un post, eliminar un post, etc.
func (s *PostService) CreatePost(
	ctx context.Context,
	post *models.Post,
) error {
	err := s.repo.Create(ctx, post)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) GetAllPosts(ctx context.Context) ([]*models.Post, error) {
	return s.repo.GetAll(ctx)
}

func (s *PostService) GetPostByID(ctx context.Context, postId uint) (*models.Post, error) {
	return s.repo.FindByID(ctx, postId)
}

func (s *PostService) GetAllPostsByUserID(ctx context.Context, userId uint) ([]*models.Post, error) {
	return s.repo.FindByUserID(ctx, userId)
}

func (s *PostService) UpdatePost(ctx context.Context, post *models.Post) error {
	return s.repo.Update(ctx, post)
}

func (s *PostService) DeletePost(ctx context.Context, postId uint) error {
	return s.repo.Delete(ctx, postId)
}
