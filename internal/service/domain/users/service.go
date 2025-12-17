package users

import (
	"context"
	"errors"

	"meu-treino-golang/users-crud/internal/service"
)

type Service struct {
	repo service.IUserRepository
}

func NewService(repo service.IUserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, name, email string) (uint, error) {
	if name == "" {
		return 0, errors.New("name cannot be empty")
	}

	if email == "" {
		return 0, errors.New("email cannot be empty")
	}

	return s.repo.Create(ctx, name, email)
}

func (s *Service) ListUsers(ctx context.Context) ([]service.UserDTO, error) {
	return s.repo.List(ctx)
}
