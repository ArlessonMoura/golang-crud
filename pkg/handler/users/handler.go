package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"meu-treino-golang/users-crud/dto"
	"meu-treino-golang/users-crud/internal/service"
)

type Handler struct {
	service service.IUserService
}

func NewHandler(svc service.IUserService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateUser(c.Request.Context(), req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) List(c *gin.Context) {
	users, err := h.service.ListUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	usersGroup := router.Group("/api/users")
	{
		usersGroup.POST("", h.Create)
		usersGroup.GET("", h.List)
	}
}
