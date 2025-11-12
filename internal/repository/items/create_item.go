package items

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

func (r *Repository) CreateItem(ctx context.Context, item *models.ItemDTO, userID uint) (uint, error) {
	if _, err := r.db.ExecContext(ctx, "SET LOCAL myapp.current_user_id = $1", userID); err != nil {
		return 0, fmt.Errorf("repository/create_item.go - failed to set local user_id: %w", err)
	}

	query := `
        INSERT INTO item (name, description, price)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	var id uint
	err := r.db.QueryRowContext(ctx, query, item.Name, item.Description, item.Price).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			switch pgErr.ConstraintName {
			case "item_name_key":
				return 0, ErrDuplicateItemName
			case "item_description_key":
				return 0, ErrDuplicateItemDescription
			}
		}
		return 0, fmt.Errorf("repository/create_item.go - failed to create item: %w", err)
	}
	return id, nil
}
