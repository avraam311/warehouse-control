package auth

import (
	"context"

	"github.com/avraam311/warehouse-control/internal/models"
	"github.com/wb-go/wbf/config"
)

type Repository interface {
	CreateUser(context.Context, *models.UserWithHashDomain) (uint, error)
	GetUser(context.Context, string) (*models.UserWithHashDB, error)
	GetUserPermissions(context.Context, string) (*models.UserPermissionsDB, error)
}

type Service struct {
	repo Repository
	cfg  *config.Config
}

func NewService(repo Repository, cfg *config.Config) *Service {
	return &Service{
		repo: repo,
		cfg:  cfg,
	}
}
