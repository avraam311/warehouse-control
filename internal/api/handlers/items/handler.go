package items

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateItem(context.Context, *models.ItemDTO) (uint, error)
	GetItems(context.Context) ([]*models.ItemDB, error)
	ReplaceItem(context.Context, uint, *models.ItemDTO) error
	DeleteItem(context.Context, uint) error
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
