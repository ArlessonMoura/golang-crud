package users

import (
	"meu-treino-golang/users-crud/internal/common"
	userService "meu-treino-golang/users-crud/internal/service/domain/users"
	userStorage "meu-treino-golang/users-crud/internal/storage/postgres/users"
)

func InitHandler(deps *common.Dependencies) *Handler {
	deps.Load()

	repo := userStorage.NewRepository(deps.DB)
	svc := userService.NewService(repo)

	return NewHandler(svc)
}
