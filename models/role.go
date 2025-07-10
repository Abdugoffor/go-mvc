package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name"`
	IsActive    bool         `json:"is_active"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}
