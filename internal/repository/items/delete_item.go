package items

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteItem(ctx context.Context, itemID uint, userID uint) error {
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to begin tx: %w", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	setUserIDQuery := fmt.Sprintf("SET LOCAL myapp.current_user_id = %d", userID)
	if _, err = tx.ExecContext(ctx, setUserIDQuery); err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to set local user_id: %w", err)
	}

	query := `
        DELETE FROM item
        WHERE id = $1;
    `

	res, err := tx.ExecContext(ctx, query, itemID)
	if err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to delete item: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository/delete_item.go - failed to commit tx: %w", err)
	}

	return nil
}
