package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
	
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (r *Repository) CreateUser(ctx context.Context, usr *models.UserWithHashDomain) (uint, error) {
	query := `
        INSERT INTO "user" (email, hash, role)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	var id uint
	err := r.db.QueryRowContext(ctx, query, usr.Email, usr.Hash, usr.Role).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			if pgErr.ConstraintName == "user_email_key" {
				return 0, ErrDuplicateEmail
			}
		}

		return 0, fmt.Errorf("repository/create_user.go - failed to create user: %w", err)
	}
	return id, nil
}
