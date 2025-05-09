package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/models"
	"github.com/onursahin/e-commerce/requests/auth"
	"github.com/onursahin/e-commerce/utils"
)

func SignIn(c *fiber.Ctx) error {
	var req auth.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials!"})
	}

	passwordIsValid := utils.CheckPasswordHash(req.Password, user.Password)

	if !passwordIsValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials!"})
	}
	return c.JSON(fiber.Map{"message": "Login successful"})
}
