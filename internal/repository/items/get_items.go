package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) GetItems(ctx context.Context) ([]*models.ItemDB, error) {
	query := `
		SELECT id, name, description, price
		FROM item
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to get items - %w", err)
	}
	defer rows.Close()

	var items []*models.ItemDB
	for rows.Next() {
		var i models.ItemDB
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Price)
		if err != nil {
			return nil, fmt.Errorf("repository/get_items.go - failed to scan item row - %w", err)
		}
		items = append(items, &i)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to iterate items rows - %w", err)
	}

	return items, nil
}
