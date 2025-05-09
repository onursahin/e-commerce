package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/models"
	"github.com/onursahin/e-commerce/requests/auth"
	"github.com/onursahin/e-commerce/utils"
)

func SignUp(c *fiber.Ctx) error {
	var request auth.SignUpRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Request do not read!"})
	}

	if err := utils.Validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": utils.FormatValidationError(err),
		})
	}

	var existingUser models.User
	err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email already exists!"})
	}

	person := models.Person{
		Name:    request.Name,
		Surname: request.Surname,
	}
	if err := database.DB.Create(&person).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Person can not created!"})
	}

	hashedPassword, err := utils.HashPassword(request.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Password could not hashed!"})
	}

	user := models.User{
		PersonID: person.ID,
		Email:    request.Email,
		Password: string(hashedPassword),
		Status:   "A",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "User can not created!"})
	}

	if err := database.DB.Preload("Person").First(&user, user.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load person data!"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created.",
		"data":    user,
	})
}
