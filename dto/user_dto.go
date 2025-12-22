// Package dto contains data transfer objects for API requests and responses.
package dto

// CreateUserRequest represents a request to create a new user.
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
