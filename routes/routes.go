package routes

import (
	"github.com/gofiber/fiber/v2"
	authHandler "github.com/onursahin/e-commerce/handlers/auth"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/healty", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "healty"})
	})

	api := app.Group("/api/v1")
	auth := api.Group("/auth")

	auth.Post("/signup", authHandler.SignUp)
	auth.Post("/login", authHandler.Login)
}
