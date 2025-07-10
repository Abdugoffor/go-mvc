package dto

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
