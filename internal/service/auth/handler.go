package auth

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"
)

type Repository interface {
	CreateUser(context.Context, *models.UserWithHashDomain) (uint, error)
	GetUser(context.Context, string) (*models.UserWithHashDB, error)
	GetUserPermissions(context.Context, string) (*models.UserPermissionsDB, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
