package users

import (
	"net/http"
	"strconv"

	"meu-treino-golang/users-crud/dto"
	"meu-treino-golang/users-crud/internal/common"
	"meu-treino-golang/users-crud/internal/service"

	"github.com/gin-gonic/gin"
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

	// convert to response DTOs
	resp := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, dto.UserResponse{ID: u.ID, Name: u.Name, Email: u.Email})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Get(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.service.GetUserByID(c.Request.Context(), uint(id64))
	if err != nil {
		if err == common.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	usersGroup := router.Group("/api/users")
	{
		usersGroup.POST("", h.Create)
		usersGroup.GET("", h.List)
		usersGroup.GET("/:id", h.Get)
	}
}
