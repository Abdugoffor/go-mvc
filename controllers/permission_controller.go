package controllers

import (
	"myapp/config"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PermissionController struct{}

func NewPermissionController() *PermissionController {
	return &PermissionController{}
}

func (c *PermissionController) PermissionIndex(ctx echo.Context) error {
	var permissions []models.Permission
	config.DB.Find(&permissions)
	return ctx.JSON(http.StatusOK, permissions)
}

func (c *PermissionController) PermissionShow(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var permission models.Permission
	result := config.DB.First(&permission, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Permission not found"})
	}
	return ctx.JSON(http.StatusOK, permission)
}

func (c *PermissionController) PermissionUpdate(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var permission models.Permission
	result := config.DB.First(&permission, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Permission not found"})
	}
	if err := ctx.Bind(&permission); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := config.DB.Save(&permission).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update permission"})
	}
	return ctx.JSON(http.StatusOK, permission)
}
