package auth

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) CreateUser(ctx context.Context, usr *models.UserWithHashDomain) (uint, error) {
	query := `
		INSERT into user
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var id uint
	err := r.db.QueryRowContext(ctx, query, usr.Email, usr.Hash, usr.Role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository/create_user.go - failed to create user - %w", err)
	}

	return id, nil
}
