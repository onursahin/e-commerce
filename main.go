package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/models"
	"github.com/onursahin/e-commerce/routes"
)

func main() {
	app := fiber.New()

	err := database.Init()
	if err != nil {
		panic("Database connection failed!")
	}

	database.DB.AutoMigrate(&models.Person{})
	database.DB.AutoMigrate(&models.User{})

	routes.RegisterRoutes(app)

	app.Listen(":8080")
}
