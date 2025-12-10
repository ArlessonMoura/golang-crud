package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "meu-treino-golang/users-crud/models"
    "meu-treino-golang/users-crud/service"
)

// UserController contém referência ao service.
type UserController struct {
    service service.UserService
}

// NewUserController cria o controller.
func NewUserController(s service.UserService) *UserController {
    return &UserController{service: s}
}

// RegisterRoutes registra as rotas no router Gin.
func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
    users := rg.Group("/users")
    {
        users.GET("", uc.GetAllUsers)
        users.POST("", uc.CreateUser)
        users.PUT("/:userId", uc.UpdateUser)
        users.DELETE("/:userId", uc.DeleteUser)
    }
}

// GetAllUsers -> GET /users
func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.service.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

// CreateUser -> POST /users
// Body JSON: { "nome": "...", "email": "..." }
func (uc *UserController) CreateUser(c *gin.Context) {
    var input struct {
        Nome  string `json:"nome" binding:"required"`
        Email string `json:"email" binding:"required,email"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := &models.User{
        Nome:  input.Nome,
        Email: input.Email,
    }

    if err := uc.service.Create(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}

// UpdateUser -> PUT /users/:userId
// Body JSON: { "nome": "...", "email": "..." }
func (uc *UserController) UpdateUser(c *gin.Context) {
    idParam := c.Param("userId")
    id64, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }
    id := uint(id64)

    // Verifica existência
    existing, err := uc.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
        return
    }
    if existing == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    var input struct {
        Nome  string `json:"nome" binding:"required"`
        Email string `json:"email" binding:"required,email"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    existing.Nome = input.Nome
    existing.Email = input.Email

    if err := uc.service.Update(existing); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
        return
    }

    c.JSON(http.StatusOK, existing)
}

// DeleteUser -> DELETE /users/:userId
func (uc *UserController) DeleteUser(c *gin.Context) {
    idParam := c.Param("userId")
    id64, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
        return
    }
    id := uint(id64)

    // Verifica existência (opcional)
    existing, err := uc.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user"})
        return
    }
    if existing == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    if err := uc.service.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}
