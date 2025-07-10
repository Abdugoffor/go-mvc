package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	CategoryController := controllers.NewCategoryController()
	ProductController := controllers.NewProductController()
	OrderController := controllers.NewOrderController()
	PermissionController := controllers.NewPermissionController()
	RoleController := controllers.NewRoleController()

	api1 := e.Group("/api/v1")

	category := api1.Group("/category")
	{
		category.GET("", CategoryController.CategoryIndex)
		category.GET("/:id", CategoryController.CategoryShow)
		category.POST("", CategoryController.CategoryStore)
		category.PUT("/:id", CategoryController.CategoryUpdate)
		category.DELETE("/:id", CategoryController.CategoryDelete)
	}

	product := api1.Group("/product")
	{
		product.GET("", ProductController.ProductIndex)
		product.GET("/:id", ProductController.ProductShow)
		product.POST("", ProductController.ProductStore)
		product.PUT("/:id", ProductController.ProductUpdate)
		product.DELETE("/:id", ProductController.ProductDelete)
	}

	order := api1.Group("/order")
	{
		order.GET("", OrderController.OrderIndex)
		order.GET("/:id", OrderController.OrderShow)
		order.POST("", OrderController.OrderStore)
		order.PUT("/:id", OrderController.OrderUpdate)
		order.DELETE("/:id", OrderController.OrderDelete)
	}

	permission := api1.Group("/permission")
	{
		permission.GET("", PermissionController.PermissionIndex)
		permission.GET("/:id", PermissionController.PermissionShow)
		permission.PUT("/:id", PermissionController.PermissionUpdate)
		// permission.POST("", PermissionController.PermissionStore)
		// permission.DELETE("/:id", PermissionController.PermissionDelete)
	}

	role := api1.Group("/role")
	{
		role.GET("", RoleController.RoleIndex)
		role.GET("/:id", RoleController.RoleShow)
		role.POST("", RoleController.RoleCreate)
		role.PUT("/:id", RoleController.RoleUpdate)
		role.DELETE("/:id", RoleController.RoleDelete)
	}

	api1.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "pong",
		})
	})
	api1.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"message": "hello",
		})
	})
}
