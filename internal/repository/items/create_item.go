package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) CreateItem(ctx context.Context, item *models.ItemDTO) (uint, error) {
	query := `
		INSERT INTO item (name, description, price)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	var id uint
	err := r.db.QueryRowContext(ctx, query, item.Name, item.Description, item.Price).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository/create_item.go - failed to create item")
	}

	return id, nil
}
