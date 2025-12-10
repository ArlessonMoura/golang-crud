package main

import (
    "github.com/gin-gonic/gin"
    "meu-treino-golang/users-crud/controller"
    "meu-treino-golang/users-crud/db"
    "meu-treino-golang/users-crud/repository"
    "meu-treino-golang/users-crud/service"
)

func main() {
    // 1. Conectar ao DB e migrar
    db.ConnectDatabase()

    // 2. Inicializar repositório, serviço e controller
    userRepo := repository.NewUserRepository()
    userSvc := service.NewUserService(userRepo)
    userCtrl := controller.NewUserController(userSvc)

    // 3. Inicializar Gin
    router := gin.Default()

    // 4. Registrar rotas (agrupando por /api, opcional)
    api := router.Group("/api")
    userCtrl.RegisterRoutes(api) // rotas ficarão em /api/users

    // 5. Iniciar servidor
    router.Run(":8080") // padrão: localhost:8080
}
