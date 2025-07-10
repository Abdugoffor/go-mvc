package dto

import "myapp/models"

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	CategoryID  int    `json:"category_id"`
}

type ProductResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Quantity     int    `json:"quantity"`
	CategoryName string `json:"category_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

func ToProductResponse(product models.Product) ProductResponse {

	var deletedAt string
	if product.DeletedAt.Valid {
		deletedAt = product.DeletedAt.Time.Format("02-01-2006 15:04:05")
	}

	return ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Quantity:     product.Quantity,
		CategoryName: product.Category.Name,
		CreatedAt:    product.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:    product.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:    deletedAt,
	}
}
