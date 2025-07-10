package controllers

import (
	"myapp/config"
	"myapp/dto"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct{}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) ProductIndex(ctx echo.Context) error {
	var products []models.Product

	config.DB.Preload("Category").Find(&products)

	var response []dto.ProductResponse
	for _, p := range products {
		response = append(response, dto.ToProductResponse(p))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) ProductShow(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var product models.Product

	result := config.DB.Preload("Category").First(&product, id)

	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	response := dto.ToProductResponse(product)
	return ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) ProductStore(ctx echo.Context) error {
	var input dto.CreateProductRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Quantity:    input.Quantity,
		CategoryID:  input.CategoryID,
	}

	config.DB.Create(&product)

	config.DB.Preload("Category").First(&product, product.ID)

	response := dto.ToProductResponse(product)

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) ProductUpdate(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	var input dto.UpdateProductRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	config.DB.Model(&product).Updates(input)

	config.DB.Preload("Category").First(&product, id)

	response := dto.ToProductResponse(product)

	return ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) ProductDelete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}
	config.DB.Delete(&product)
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Product deleted"})
}
