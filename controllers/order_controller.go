package controllers

import (
	"myapp/config"
	"myapp/dto"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct{}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) OrderIndex(ctx echo.Context) error {
	var orders []models.Order
	config.DB.Preload("Product").Find(&orders)

	var response []dto.OrderResponse
	for _, or := range orders {
		response = append(response, dto.ToOrderResponse(or))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) OrderShow(ctx echo.Context) error {

	id, _ := strconv.Atoi(ctx.Param("id"))

	var order models.Order

	result := config.DB.Preload("Product").First(&order, id)

	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Order not found"})
	}

	response := dto.ToOrderResponse(order)
	return ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) OrderStore(ctx echo.Context) error {
	var input dto.CreateOrderRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var product models.Product

	result := config.DB.First(&product, input.ProductID)

	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}
	amount := product.Price * input.Quantity

	order := models.Order{
		ProductId: input.ProductID,
		Quantity:  input.Quantity,
		Amount:    amount,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create order"})
	}

	config.DB.Preload("Product").First(&order, order.ID)

	response := dto.ToOrderResponse(order)
	return ctx.JSON(http.StatusCreated, response)
}

func (c *OrderController) OrderUpdate(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Order not found"})
	}

	var input dto.UpdateOrderRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := ctx.Validate(input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var product models.Product

	result := config.DB.First(&product, input.ProductID)

	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	amount := product.Price * input.Quantity

	orderUp := models.Order{
		ProductId: input.ProductID,
		Quantity:  input.Quantity,
		Amount:    amount,
	}

	config.DB.Model(&order).Updates(orderUp)

	config.DB.Preload("Product").First(&order, id)

	response := dto.ToOrderResponse(order)

	return ctx.JSON(http.StatusOK, response)
}

func (c *OrderController) OrderDelete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Order not found"})
	}
	config.DB.Delete(&order)
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Order deleted"})
}
