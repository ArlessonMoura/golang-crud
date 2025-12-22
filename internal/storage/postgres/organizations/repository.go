// Package organizations provides database access for organization data.
package organizations

import (
	"meu-treino-golang/users-crud/dto"

	"gorm.io/gorm"
)

type OrganizationModel struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"not null"`
	Users []OrgUserModel
}

type OrgUserModel struct {
	ID         uint   `gorm:"primaryKey"`
	OrgID      uint   `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
	Permission string `gorm:"not null;default:'READ'"`

	Organization OrganizationModel `gorm:"foreignKey:OrgID;constraint:OnDelete:CASCADE"`
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// CreateOrg creates a new organization.
func (r *Repository) CreateOrg(orgName string) (uint, error) {
	org := OrganizationModel{Name: orgName}
	if err := r.db.Create(&org).Error; err != nil {
		return 0, err
	}
	return org.ID, nil
}

func (r *Repository) GetOrg(orgID uint) (*OrganizationModel, error) {
	var org OrganizationModel
	if err := r.db.Preload("Users").First(&org, orgID).Error; err != nil {
		return nil, err
	}
	return &org, nil
}

func (r *Repository) ListOrgs() ([]OrganizationModel, error) {
	var orgs []OrganizationModel
	if err := r.db.Find(&orgs).Error; err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r *Repository) UpdateOrg(orgID uint, name string) error {
	return r.db.Model(&OrganizationModel{}).Where("id = ?", orgID).Update("name", name).Error
}

func (r *Repository) DeleteOrg(orgID uint) error {
	return r.db.Delete(&OrganizationModel{}, orgID).Error
}

// AddUserToOrg adds a user to an organization.
func (r *Repository) AddUserToOrg(orgID, userID uint, permission dto.PermissionType) error {
	orgUser := OrgUserModel{
		OrgID:      orgID,
		UserID:     userID,
		Permission: string(permission),
	}
	return r.db.Create(&orgUser).Error
}

func (r *Repository) GetOrgUsers(orgID uint) ([]OrgUserModel, error) {
	var users []OrgUserModel
	if err := r.db.Where("org_id = ?", orgID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) UpdateUserPermission(orgID, userID uint, permission dto.PermissionType) error {
	return r.db.Model(&OrgUserModel{}).
		Where("org_id = ? AND user_id = ?", orgID, userID).
		Update("permission", string(permission)).
		Error
}

func (r *Repository) RemoveUserFromOrg(orgID, userID uint) error {
	return r.db.Delete(&OrgUserModel{}, "org_id = ? AND user_id = ?", orgID, userID).Error
}

func (r *Repository) GetUserPermissionInOrg(orgID, userID uint) (dto.PermissionType, error) {
	var user OrgUserModel
	if err := r.db.Where("org_id = ? AND user_id = ?", orgID, userID).First(&user).Error; err != nil {
		return "", err
	}
	return dto.PermissionType(user.Permission), nil
}
