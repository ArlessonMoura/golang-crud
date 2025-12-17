package service

import "context"

type IUserService interface {
	CreateUser(ctx context.Context, name, email string) (uint, error)
	ListUsers(ctx context.Context) ([]UserDTO, error)
}

type UserDTO struct {
	ID    uint
	Name  string
	Email string
}
