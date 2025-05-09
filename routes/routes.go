package routes

import (
	"github.com/gofiber/fiber/v2"
	authHandler "github.com/onursahin/e-commerce/handlers/auth"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/healty", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "healty"})
	})

	auth := app.Group("/auth")
	auth.Post("/signup", authHandler.SignUp)
	auth.Post("/signin", authHandler.SignIn)
}
