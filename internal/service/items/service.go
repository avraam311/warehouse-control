package items

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"
)

type Repository interface {
	CreateItem(context.Context, *models.ItemDTO, uint) (uint, error)
	GetItems(context.Context, uint) ([]*models.ItemDB, error)
	ReplaceItem(context.Context, uint, *models.ItemDTO, uint) error
	DeleteItem(context.Context, uint, uint) error
	GetHistory(context.Context, *models.HistoryFilter, uint) ([]*models.HistoryItem, error)
	ExportHistoryCSV(context.Context, *models.HistoryFilter, uint) ([]byte, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
