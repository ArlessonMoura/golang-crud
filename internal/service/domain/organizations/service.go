// Package organizations provides business logic for organizations.
package organizations

import (
	"context"
	"errors"

	"meu-treino-golang/users-crud/dto"
	"meu-treino-golang/users-crud/internal/storage/postgres/organizations"
)

type IOrganizationService interface {
	CreateOrg(ctx context.Context, name string) (uint, error)
	GetOrg(ctx context.Context, orgID uint) (*OrganizationDTO, error)
	ListOrgs(ctx context.Context) ([]OrganizationDTO, error)
	UpdateOrg(ctx context.Context, orgID uint, name string) error
	DeleteOrg(ctx context.Context, orgID uint) error
	
	AddUserToOrg(ctx context.Context, orgID, userID uint, permission dto.PermissionType) error
	GetOrgUsers(ctx context.Context, orgID uint) ([]OrgUserDTO, error)
	UpdateUserPermission(ctx context.Context, orgID, userID uint, permission dto.PermissionType) error
	RemoveUserFromOrg(ctx context.Context, orgID, userID uint) error
	GetUserPermissionInOrg(ctx context.Context, orgID, userID uint) (dto.PermissionType, error)
}

type OrganizationDTO struct {
	ID   uint
	Name string
}

type OrgUserDTO struct {
	UserID     uint
	OrgID      uint
	Permission dto.PermissionType
}

type Service struct {
	repo *organizations.Repository
}

func NewService(repo *organizations.Repository) *Service {
	return &Service{repo: repo}
}

// CreateOrg creates a new organization.
func (s *Service) CreateOrg(ctx context.Context, name string) (uint, error) {
	if name == "" {
		return 0, errors.New("organization name cannot be empty")
	}
	return s.repo.CreateOrg(name)
}

func (s *Service) GetOrg(ctx context.Context, orgID uint) (*OrganizationDTO, error) {
	org, err := s.repo.GetOrg(orgID)
	if err != nil {
		return nil, err
	}
	return &OrganizationDTO{
		ID:   org.ID,
		Name: org.Name,
	}, nil
}

func (s *Service) ListOrgs(ctx context.Context) ([]OrganizationDTO, error) {
	orgs, err := s.repo.ListOrgs()
	if err != nil {
		return nil, err
	}
	
	dtos := make([]OrganizationDTO, 0, len(orgs))
	for _, org := range orgs {
		dtos = append(dtos, OrganizationDTO{
			ID:   org.ID,
			Name: org.Name,
		})
	}
	return dtos, nil
}

func (s *Service) UpdateOrg(ctx context.Context, orgID uint, name string) error {
	if name == "" {
		return errors.New("organization name cannot be empty")
	}
	return s.repo.UpdateOrg(orgID, name)
}

func (s *Service) DeleteOrg(ctx context.Context, orgID uint) error {
	return s.repo.DeleteOrg(orgID)
}

// AddUserToOrg adds a user to an organization.
func (s *Service) AddUserToOrg(ctx context.Context, orgID, userID uint, permission dto.PermissionType) error {
	if !isValidPermission(permission) {
		return errors.New("invalid permission type")
	}
	return s.repo.AddUserToOrg(orgID, userID, permission)
}

func (s *Service) GetOrgUsers(ctx context.Context, orgID uint) ([]OrgUserDTO, error) {
	users, err := s.repo.GetOrgUsers(orgID)
	if err != nil {
		return nil, err
	}
	
	dtos := make([]OrgUserDTO, 0, len(users))
	for _, user := range users {
		dtos = append(dtos, OrgUserDTO{
			UserID:     user.UserID,
			OrgID:      user.OrgID,
			Permission: dto.PermissionType(user.Permission),
		})
	}
	return dtos, nil
}

func (s *Service) UpdateUserPermission(ctx context.Context, orgID, userID uint, permission dto.PermissionType) error {
	if !isValidPermission(permission) {
		return errors.New("invalid permission type")
	}
	return s.repo.UpdateUserPermission(orgID, userID, permission)
}

func (s *Service) RemoveUserFromOrg(ctx context.Context, orgID, userID uint) error {
	return s.repo.RemoveUserFromOrg(orgID, userID)
}

func (s *Service) GetUserPermissionInOrg(ctx context.Context, orgID, userID uint) (dto.PermissionType, error) {
	return s.repo.GetUserPermissionInOrg(orgID, userID)
}

func isValidPermission(permission dto.PermissionType) bool {
	return permission == dto.PermissionRead || 
		   permission == dto.PermissionWrite || 
		   permission == dto.PermissionRoot
}
