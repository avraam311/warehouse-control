package auth

import (
	"errors"

	"github.com/wb-go/wbf/dbpg"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Repository struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
