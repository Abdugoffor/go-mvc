package controllers

import (
	"myapp/config"
	"myapp/dto"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoleController struct{}

func NewRoleController() *RoleController {
	return &RoleController{}
}

func (c *RoleController) RoleIndex(ctx echo.Context) error {
	var roles []models.Role
	if err := config.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch roles"})
	}

	var response []dto.RoleResponse
	for _, r := range roles {
		response = append(response, dto.ToRoleResponse(r))
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c *RoleController) RoleShow(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var role models.Role

	result := config.DB.Preload("Permissions").First(&role, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Role not found"})
	}
	response := dto.ToRoleResponse(role)
	return ctx.JSON(http.StatusOK, response)
}

func (c *RoleController) RoleCreate(ctx echo.Context) error {
	var role models.Role
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := config.DB.Save(&role).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create role"})
	}
	return ctx.JSON(http.StatusCreated, role)
}

func (c *RoleController) RoleUpdate(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var role models.Role
	result := config.DB.First(&role, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Role not found"})
	}
	if err := ctx.Bind(&role); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := config.DB.Save(&role).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update role"})
	}
	return ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) RoleDelete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var role models.Role
	result := config.DB.First(&role, id)
	if result.Error != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Role not found"})
	}
	if err := config.DB.Delete(&role).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete role"})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "Role deleted"})
}
