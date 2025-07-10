package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductId int `json:"product_id"`
	Product   Product
	Quantity  int `json:"quantity"`
	Amount    int `json:"amount"`
}

func (Order) TableName() string {
	return "orders"
}
