package db

import (
    "log"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "meu-treino-golang/users-crud/models"
)

var DB *gorm.DB

// ConnectDatabase abre conexão com SQLite e executa AutoMigrate.
func ConnectDatabase() {
    database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migra o modelo User (cria tabela se não existir)
    err = database.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    DB = database
}
