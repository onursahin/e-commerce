package utils

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func RespondSuccess(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func RespondError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Success: false,
		Message: message,
	})
}

func RespondValidationError(c *fiber.Ctx, statusCode int, errors interface{}) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Success: false,
		Errors:  errors,
	})
}
