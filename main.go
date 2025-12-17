package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"meu-treino-golang/users-crud/internal/common"
	"meu-treino-golang/users-crud/internal/storage/postgres/users"
	"meu-treino-golang/users-crud/routes"
)

func main() {
	// 1. Conectar ao banco de dados
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 2. AutoMigrate UserModel
	if err := database.AutoMigrate(&users.UserModel{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 3. Inicializar dependências
	deps := &common.Dependencies{
		DB: database,
	}

	// 4. Inicializar Gin
	router := gin.Default()

	// 5. Registrar rotas
	routes.RegisterRoutes(router, deps)

	// 6. Iniciar servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
