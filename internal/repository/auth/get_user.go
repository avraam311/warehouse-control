package auth

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) GetUser(ctx context.Context, email string) (*models.UserWithHashDB, error) {
	query := `
		SELECT id, hash, role
		FROM user
		WHERE email = $1;
	`

	user := models.UserWithHashDB{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Hash)
	if err != nil {
		return nil, fmt.Errorf("repository/get_user.go - failed to get user - %w", err)
	}

	return &user, nil
}
