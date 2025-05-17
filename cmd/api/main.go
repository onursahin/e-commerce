package main

import (
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/models"
	"github.com/onursahin/e-commerce/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New())

	err := database.Init()
	if err != nil {
		panic("Database connection failed!")
	}

	database.DB.AutoMigrate(&models.Person{})
	database.DB.AutoMigrate(&models.User{})

	routes.RegisterRoutes(app)

	app.Listen(":8080")
}
