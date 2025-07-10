package dto

import "myapp/models"

type CreateOrderRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type UpdateOrderRequest struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderResponse struct {
	ID        uint   `json:"id"`
	Product   string `json:"product"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func ToOrderResponse(order models.Order) OrderResponse {
	var deletedAt string
	if order.DeletedAt.Valid {
		deletedAt = order.DeletedAt.Time.Format("02-01-2006 15:04:05")
	}

	return OrderResponse{
		ID:        order.ID,
		Product:   order.Product.Name,
		Price:     order.Product.Price,
		Quantity:  order.Quantity,
		Amount:    order.Amount,
		CreatedAt: order.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt: deletedAt,
	}
}
