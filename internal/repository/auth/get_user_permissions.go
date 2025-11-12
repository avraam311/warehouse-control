package auth

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) GetUserPermissions(ctx context.Context, email string) (*models.UserPermissionsDB, error) {
	query := `
		SELECT routes
		FROM user_permission
		WHERE email = $1;
	`

	permissions := models.UserPermissionsDB{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(&permissions.Routes)
	if err != nil {
		return nil, fmt.Errorf("repository/get_user_permissions.go - failed to get user permissions - %w", err)
	}

	return &permissions, nil
}
