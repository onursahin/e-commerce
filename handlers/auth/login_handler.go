package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/repositories"
	"github.com/onursahin/e-commerce/requests/auth"
	"github.com/onursahin/e-commerce/utils"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var request auth.LoginRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.RespondError(c, fiber.StatusBadRequest, "Invalid request")
	}

	if err := utils.Validate.Struct(&request); err != nil {
		errors := utils.FormatValidationError(err, &request)
		return utils.RespondValidationError(c, fiber.StatusBadRequest, errors)
	}

	userRepository := repositories.NewUserRepository(database.DB)

	user, err := userRepository.GetOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("email = ?", request.Email)
	})

	if err != nil {
		return utils.RespondError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		return utils.RespondError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := utils.GenerateToken(user.Email, int64(user.ID))
	if err != nil {
		return utils.RespondError(c, fiber.StatusInternalServerError, "Could not authenticate user")
	}

	return utils.RespondSuccess(c, fiber.StatusOK, fiber.Map{
		"message": "User sign in successfully.",
		"token":   token,
	})
}
