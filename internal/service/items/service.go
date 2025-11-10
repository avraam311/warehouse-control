package items

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"
)

type Repository interface {
	CreateItem(context.Context, *models.ItemDTO) (uint, error)
	GetItems(context.Context) ([]*models.ItemDB, error)
	ReplaceItem(context.Context, uint, *models.ItemDTO) error
	DeleteItem(context.Context, uint) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
