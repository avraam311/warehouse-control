package auth

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	Register(context.Context, *models.UserWithRoleDTO) (uint, error)
	Login(context.Context, *models.UserDTO) (string, error)
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
