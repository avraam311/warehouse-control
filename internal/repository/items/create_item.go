package items

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"

	"github.com/lib/pq"
)

func (r *Repository) CreateItem(ctx context.Context, item *models.ItemDTO, userID uint) (uint, error) {
	query := `
        INSERT INTO item (name, description, price)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	var id uint
	err := r.db.QueryRowContext(ctx, query, item.Name, item.Description, item.Price).Scan(&id)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				switch pqErr.Constraint {
				case "item_name_key":
					return 0, ErrDuplicateItemName
				case "item_description_key":
					return 0, ErrDuplicateItemDescription
				}
			}
		}
		return 0, fmt.Errorf("repository/create_item.go - failed to create item: %w", err)
	}
	return id, nil
}
