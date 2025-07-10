package main

import (
	"myapp/config"
	"myapp/routes"
	"myapp/seed"
	"myapp/validator"

	validatorV10 "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()

	e := echo.New()

	e.Validator = &validator.CustomValidator{Validator: validatorV10.New()}

	routes.RegisterRoutes(e)

	seed.SeedPermissions(config.DB, e)

	e.Logger.Fatal(e.Start(":8080"))
}
