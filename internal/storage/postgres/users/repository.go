// Package users provides database access for user data.
package users

import (
	"context"

	"meu-treino-golang/users-crud/internal/service"

	"gorm.io/gorm"
)

type UserModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Email string `gorm:"uniqueIndex;not null"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, name, email string) (uint, error) {
	user := UserModel{Name: name, Email: email}
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *Repository) List(ctx context.Context) ([]service.UserDTO, error) {
	var models []UserModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]service.UserDTO, 0, len(models))
	for _, m := range models {
		users = append(users, service.UserDTO{
			ID:    m.ID,
			Name:  m.Name,
			Email: m.Email,
		})
	}

	return users, nil
}
func (r *Repository) GetByID(ctx context.Context, id uint) (*UserModel, error) {
	var user UserModel
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}