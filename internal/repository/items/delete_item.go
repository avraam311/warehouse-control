package items

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteItem(ctx context.Context, itemID uint, userID uint) error {
	if _, err := r.db.ExecContext(ctx, "SET LOCAL myapp.current_user_id = $1", userID); err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to set local user_id: %w", err)
	}

	query := `
        DELETE FROM item
        WHERE id = $1;
    `

	res, err := r.db.ExecContext(ctx, query, itemID)
	if err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to delete item: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	return nil
}
