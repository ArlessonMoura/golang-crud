package users

import "meu-treino-golang/users-crud/internal/service"

var _ service.IUserRepository = (*Repository)(nil)
