package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/onursahin/e-commerce/database"
	"github.com/onursahin/e-commerce/models"
	"github.com/onursahin/e-commerce/repositories"
	"github.com/onursahin/e-commerce/requests/auth"
	"github.com/onursahin/e-commerce/utils"
	"gorm.io/gorm"
)

func SignUp(c *fiber.Ctx) error {
	var request auth.SignUpRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.RespondError(c, fiber.StatusBadRequest, "Invalid request")
	}

	if err := utils.Validate.Struct(&request); err != nil {
		errors := utils.FormatValidationError(err, &request)
		return utils.RespondValidationError(c, fiber.StatusBadRequest, errors)
	}

	userRepository := repositories.NewUserRepository(database.DB)
	personRepository := repositories.NewPersonRepository(database.DB)

	_, err := userRepository.GetOne(func(db *gorm.DB) *gorm.DB {
		return db.Where("email = ?", request.Email)
	})

	if err == nil {
		return utils.RespondError(c, fiber.StatusBadRequest, "Email already exists")
	}

	person := models.Person{
		Name:    request.Name,
		Surname: request.Surname,
	}

	if err := personRepository.Create(&person); err != nil {
		return utils.RespondError(c, fiber.StatusInternalServerError, "Person could not be created")
	}

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return utils.RespondError(c, fiber.StatusInternalServerError, "Password could not be hashed")
	}

	user := models.User{
		PersonID: person.ID,
		Email:    request.Email,
		Password: hashedPassword,
	}
	if err := userRepository.Create(&user); err != nil {
		return utils.RespondError(c, fiber.StatusInternalServerError, "User could not be created")
	}

	createdUser, err := userRepository.GetByID(user.ID, "Person")
	if err != nil {
		return utils.RespondError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	return utils.RespondSuccess(c, fiber.StatusCreated, fiber.Map{
		"message": "User created.",
		"user":    createdUser,
	})
}
