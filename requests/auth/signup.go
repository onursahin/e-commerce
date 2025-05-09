package auth

import (
	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *SignUpRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
} 