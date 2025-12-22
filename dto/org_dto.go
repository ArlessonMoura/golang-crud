// Package dto contains data transfer objects for API requests and responses.
package dto

// CreateOrganizationRequest represents a request to create a new organization.
type CreateOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateOrganizationRequest struct {
	Name string `json:"name" binding:"required"`
}

type OrganizationResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// PermissionType represents user permissions in an organization.
type PermissionType string

const (
	PermissionRead  PermissionType = "READ"
	PermissionWrite PermissionType = "WRITE"
	PermissionRoot  PermissionType = "ROOT"
)

// AddUserToOrgRequest represents a request to add a user to an organization.
type AddUserToOrgRequest struct {
	UserID     uint           `json:"user_id" binding:"required"`
	Permission PermissionType `json:"permission" binding:"required"`
}

type UpdateOrgUserPermissionRequest struct {
	Permission PermissionType `json:"permission" binding:"required"`
}

type OrgUserResponse struct {
	UserID     uint           `json:"user_id"`
	UserName   string         `json:"user_name"`
	UserEmail  string         `json:"user_email"`
	OrgID      uint           `json:"org_id"`
	Permission PermissionType `json:"permission"`
}

type OrganizationDetailResponse struct {
	ID    uint               `json:"id"`
	Name  string             `json:"name"`
	Users []OrgUserResponse  `json:"users"`
}
