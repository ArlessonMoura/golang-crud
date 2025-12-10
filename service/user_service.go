package service

import (
    "meu-treino-golang/users-crud/models"
    "meu-treino-golang/users-crud/repository"
)

// UserService define os métodos do serviço.
type UserService interface {
    Create(user *models.User) error
    GetAll() ([]models.User, error)
    GetByID(id uint) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
}

type userService struct {
    repo repository.UserRepository
}

// NewUserService cria uma instância do serviço.
func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) Create(user *models.User) error {
    // Aqui você poderia adicionar validações de negócio (ex.: checar formato do email).
    return s.repo.Create(user)
}

func (s *userService) GetAll() ([]models.User, error) {
    return s.repo.GetAll()
}

func (s *userService) GetByID(id uint) (*models.User, error) {
    return s.repo.GetByID(id)
}

func (s *userService) Update(user *models.User) error {
    return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
    return s.repo.Delete(id)
}
