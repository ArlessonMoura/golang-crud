package routes

import (
	"github.com/gin-gonic/gin"
	"meu-treino-golang/users-crud/internal/common"
	usersHandler "meu-treino-golang/users-crud/pkg/handler/users"
)

func RegisterRoutes(router *gin.Engine, deps *common.Dependencies) {
	handler := usersHandler.InitHandler(deps)
	handler.RegisterRoutes(router)
}
