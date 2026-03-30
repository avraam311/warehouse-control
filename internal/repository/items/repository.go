package items

import (
	"context"
	"errors"

	"github.com/avraam311/warehouse-control/internal/models"
	"github.com/wb-go/wbf/dbpg"
)

var (
	ErrItemNotFound             = errors.New("item not found")
	ErrDuplicateItemName        = errors.New("item name already exists")
	ErrDuplicateItemDescription = errors.New("item description already exists")
)

type Repository struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *Repository {
	return &Repository{
		db: db,
	}
}

type RepositoryI interface {
	CreateItem(context.Context, *models.ItemDTO, uint) (uint, error)
	GetItems(context.Context, uint) ([]*models.ItemDB, error)
	ReplaceItem(context.Context, uint, *models.ItemDTO, uint) error
	DeleteItem(context.Context, uint, uint) error
	GetHistory(context.Context, *models.HistoryFilter, uint) ([]*models.HistoryItem, error)
	ExportHistoryCSV(context.Context, *models.HistoryFilter, uint) ([]byte, error)
}
