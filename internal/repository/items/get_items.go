package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) GetItems(ctx context.Context, userID uint) ([]*models.ItemDB, error) {
	var items []*models.ItemDB

	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	setUserIDQuery := fmt.Sprintf("SET LOCAL myapp.current_user_id = %d", userID)
	if _, err = tx.ExecContext(ctx, setUserIDQuery); err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to set local user_id: %w", err)
	}

	query := `
        SELECT id, name, description, price
        FROM item
    `

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to get items - %w", err)
	}
	defer rows.Close()

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

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("repository/get_items.go - failed to commit tx: %w", err)
	}

	return items, nil
}
