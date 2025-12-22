// Package organizations handles all HTTP requests for organization management.
package organizations

import (
	"net/http"
	"strconv"

	"meu-treino-golang/users-crud/dto"
	orgService "meu-treino-golang/users-crud/internal/service/domain/organizations"
	usersStorage "meu-treino-golang/users-crud/internal/storage/postgres/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	orgService orgService.IOrganizationService
	db         *gorm.DB
}

func NewHandler(service orgService.IOrganizationService, db *gorm.DB) *Handler {
	return &Handler{
		orgService: service,
		db:         db,
	}
}

// CreateOrg creates a new organization.
func (h *Handler) CreateOrg(c *gin.Context) {
	var req dto.CreateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.orgService.CreateOrg(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) ListOrgs(c *gin.Context) {
	orgs, err := h.orgService.ListOrgs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.OrganizationResponse, 0, len(orgs))
	for _, org := range orgs {
		response = append(response, dto.OrganizationResponse{
			ID:   org.ID,
			Name: org.Name,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetOrg(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	org, err := h.orgService.GetOrg(c.Request.Context(), uint(orgID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "organization not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get organization users
	users, err := h.orgService.GetOrgUsers(c.Request.Context(), uint(orgID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch user details for each org user
	userResponses := make([]dto.OrgUserResponse, 0, len(users))
	usersRepo := usersStorage.NewRepository(h.db)
	for _, user := range users {
		userModel, err := usersRepo.GetByID(c.Request.Context(), user.UserID)
		if err == nil && userModel != nil {
			userResponses = append(userResponses, dto.OrgUserResponse{
				UserID:     user.UserID,
				UserName:   userModel.Name,
				UserEmail:  userModel.Email,
				OrgID:      user.OrgID,
				Permission: user.Permission,
			})
		}
	}

	response := dto.OrganizationDetailResponse{
		ID:    org.ID,
		Name:  org.Name,
		Users: userResponses,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateOrg(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	var req dto.UpdateOrganizationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user has WRITE permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionWrite, dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	if err := h.orgService.UpdateOrg(c.Request.Context(), uint(orgID), req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "organization updated successfully"})
}

func (h *Handler) DeleteOrg(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	// Check if user has ROOT permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	if err := h.orgService.DeleteOrg(c.Request.Context(), uint(orgID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "organization deleted successfully"})
}

// AddUserToOrg adds a user to an organization.
func (h *Handler) AddUserToOrg(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	// Check if user has ROOT permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	var req dto.AddUserToOrgRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orgService.AddUserToOrg(c.Request.Context(), uint(orgID), req.UserID, req.Permission); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user added to organization"})
}

func (h *Handler) ListOrgUsers(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	// Check if user has READ permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionRead, dto.PermissionWrite, dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	users, err := h.orgService.GetOrgUsers(c.Request.Context(), uint(orgID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch user details for each org user
	usersRepo := usersStorage.NewRepository(h.db)
	response := make([]dto.OrgUserResponse, 0, len(users))
	for _, user := range users {
		userModel, err := usersRepo.GetByID(c.Request.Context(), user.UserID)
		if err == nil && userModel != nil {
			response = append(response, dto.OrgUserResponse{
				UserID:     user.UserID,
				UserName:   userModel.Name,
				UserEmail:  userModel.Email,
				OrgID:      user.OrgID,
				Permission: user.Permission,
			})
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateUserPermission(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Check if user has ROOT permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	var req dto.UpdateOrgUserPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orgService.UpdateUserPermission(c.Request.Context(), uint(orgID), uint(userID), req.Permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user permission updated successfully"})
}

func (h *Handler) RemoveUserFromOrg(c *gin.Context) {
	orgID, err := strconv.ParseUint(c.Param("orgId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid organization id"})
		return
	}

	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Check if user has ROOT permission
	if !h.hasOrgPermission(c, uint(orgID), []dto.PermissionType{dto.PermissionRoot}) {
		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		return
	}

	if err := h.orgService.RemoveUserFromOrg(c.Request.Context(), uint(orgID), uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user removed from organization"})
}

// Helper methods
func (h *Handler) hasOrgPermission(c *gin.Context, orgID uint, requiredPermissions []dto.PermissionType) bool {
	// TODO: Extract user ID from context/token
	// For now, this is a placeholder that should be implemented with proper authentication
	// In a real scenario, you would get the user ID from JWT token or session
	userID, exists := c.Get("userID")
	if !exists {
		// Allow all operations if no user context (for development)
		// In production, this should return false
		return true
	}

	permission, err := h.orgService.GetUserPermissionInOrg(c.Request.Context(), orgID, userID.(uint))
	if err != nil {
		return false
	}

	for _, required := range requiredPermissions {
		if permission == required {
			return true
		}
	}
	return false
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	{
		// Organizations
		orgGroup := apiGroup.Group("/org")
		{
			orgGroup.POST("", h.CreateOrg)
			orgGroup.GET("", h.ListOrgs)
			orgGroup.GET("/:orgId", h.GetOrg)
			orgGroup.PUT("/:orgId", h.UpdateOrg)
			orgGroup.DELETE("/:orgId", h.DeleteOrg)

			// Organization Users
			usersGroup := orgGroup.Group("/:orgId/users")
			{
				usersGroup.POST("", h.AddUserToOrg)
				usersGroup.GET("", h.ListOrgUsers)
				usersGroup.PUT("/:userId", h.UpdateUserPermission)
				usersGroup.DELETE("/:userId", h.RemoveUserFromOrg)
			}
		}
	}
}
