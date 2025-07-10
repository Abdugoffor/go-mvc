package models

import "gorm.io/gorm"

type PermissionGroup struct {
	gorm.Model
	Name        string       `json:"name"`
	IsActive    bool         `json:"is_active"`
	Permissions []Permission `gorm:"foreignKey:PermissionGroupID"`
}

func (PermissionGroup) TableName() string {
	return "permission_groups"
}
