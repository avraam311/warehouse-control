package items

import (
	"errors"

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
