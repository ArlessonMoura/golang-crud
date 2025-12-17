package service

import "context"

type IUserRepository interface {
	Create(ctx context.Context, name, email string) (uint, error)
	List(ctx context.Context) ([]UserDTO, error)
}
