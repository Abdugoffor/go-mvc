package controllers

import (
	"myapp/config"
	"myapp/dto"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (c *CategoryController) CategoryIndex(ctx echo.Context) error {
	var categories []models.Category
	config.DB.Where("is_active = ?", true).Find(&categories)
	return ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) CategoryShow(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var category models.Category
	result := config.DB.Where("is_active = ?", true).First(&category, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Category not found"})
	}
	return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) CategoryStore(ctx echo.Context) error {
	var input dto.CreateCategoryRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
		IsActive:    input.IsActive,
	}

	config.DB.Create(&category)
	return ctx.JSON(http.StatusCreated, category)
}

func (c *CategoryController) CategoryUpdate(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Category not found"})
	}

	var input dto.UpdateCategoryRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	config.DB.Model(&category).Updates(input)
	return ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) CategoryDelete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var category models.Category
	if err := config.DB.First(&category, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Category not found"})
	}
	config.DB.Delete(&category)
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Category deleted"})
}
