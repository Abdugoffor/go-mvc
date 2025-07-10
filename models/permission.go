package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name              string `json:"name"`
	Key               string `json:"key"` // GET:/api/v1/category
	IsActive          bool   `json:"is_active"`
	PermissionGroupID uint
}

func (Permission) TableName() string {
	return "permissions"
}
