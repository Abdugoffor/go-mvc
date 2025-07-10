package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

func (Category) TableName() string {
	return "categories"
}
