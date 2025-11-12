package auth

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Register(ctx context.Context, user *models.UserWithRoleDTO) (uint, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("service/register.go - failed to hash password - %w", err)
	}
	userWithHash := models.UserWithHashDomain{
		Email: user.Email,
		Hash:  hash,
		Role:  user.Role,
	}

	id, err := s.repo.CreateUser(ctx, &userWithHash)
	if err != nil {
		return 0, fmt.Errorf("service/register.go - %w", err)
	}

	return id, nil
}
