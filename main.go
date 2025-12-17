package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"meu-treino-golang/users-crud/internal/common"
	"meu-treino-golang/users-crud/internal/storage/postgres/users"
	"meu-treino-golang/users-crud/routes"
)

func main() {
	// 1. Conectar ao banco de dados PostgreSQL
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default DSN para desenvolvimento local
		dsn = "host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=UTC"
		log.Println("DATABASE_URL not set. Using default DSN for local development.")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database:", err)
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
