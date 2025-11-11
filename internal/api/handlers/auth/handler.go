package auth

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	Register(context.Context, *models.UserDTOWithRole) (uint, error)
	Login(context.Context, *models.UserDTO) (*jwt.Token, error)
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
