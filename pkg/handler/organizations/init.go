// Package organizations handles HTTP requests for organization management.
package organizations

import (
	"meu-treino-golang/users-crud/internal/common"
	orgService "meu-treino-golang/users-crud/internal/service/domain/organizations"
	orgStorage "meu-treino-golang/users-crud/internal/storage/postgres/organizations"
)

func InitHandler(deps *common.Dependencies) *Handler {
	deps.Load()

	repo := orgStorage.NewRepository(deps.DB)
	service := orgService.NewService(repo)

	return NewHandler(service, deps.DB)
}
