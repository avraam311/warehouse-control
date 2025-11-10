package items

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteItem(ctx context.Context, id uint) error {
	query := `
        DELETE FROM item
        WHERE id = $1;
    `

	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to delete item: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	return nil
}
