package service

import "context"

type IUserService interface {
	CreateUser(ctx context.Context, name, email string) (uint, error)
	ListUsers(ctx context.Context) ([]UserDTO, error)
	GetUserByID(ctx context.Context, id uint) (*UserDTO, error)
}

type UserDTO struct {
	ID    uint
	Name  string
	Email string
}
