package repository

import (
    "errors"

    "meu-treino-golang/users-crud/db"
    "meu-treino-golang/users-crud/models"
    "gorm.io/gorm"
)


// UserRepository define os métodos do repositório.
type UserRepository interface {
    Create(user *models.User) error
    GetAll() ([]models.User, error)
    GetByID(id uint) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
}

// userRepository é a implementação concreta.
type userRepository struct {
    db *gorm.DB
}

// NewUserRepository cria uma instância de UserRepository.
func NewUserRepository() UserRepository {
    return &userRepository{
        db: db.DB,
    }
}

func (r *userRepository) Create(user *models.User) error {
    result := r.db.Create(user) // ID será preenchido automaticamente
    return result.Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
    var users []models.User
    result := r.db.Find(&users)
    return users, result.Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
    var user models.User
    result := r.db.First(&user, uint(id))
    if errors.Is(result.Error, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, result.Error
}

func (r *userRepository) Update(user *models.User) error {
    // Assegura que o usuário existe antes de atualizar (opcional)
    result := r.db.Save(user)
    return result.Error
}

func (r *userRepository) Delete(id uint) error {
    result := r.db.Delete(&models.User{}, id)
    return result.Error
}
