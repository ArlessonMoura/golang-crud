// Package routes registers all HTTP routes for the application.
package routes

import (
	"meu-treino-golang/users-crud/internal/common"
	orgHandler "meu-treino-golang/users-crud/pkg/handler/organizations"
	usersHandler "meu-treino-golang/users-crud/pkg/handler/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *common.Dependencies) {
	usersHandlerInstance := usersHandler.InitHandler(deps)
	usersHandlerInstance.RegisterRoutes(router)

	orgsHandlerInstance := orgHandler.InitHandler(deps)
	orgsHandlerInstance.RegisterRoutes(router)
}
