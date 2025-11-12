package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) ReplaceItem(ctx context.Context, itemID uint, item *models.ItemDTO, userID uint) error {
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	setUserIDQuery := fmt.Sprintf("SET LOCAL myapp.current_user_id = %d", userID)
	if _, err = tx.ExecContext(ctx, setUserIDQuery); err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to set local user_id: %w", err)
	}

	query := `
        UPDATE item
        SET name = $1, description = $2, price = $3
        WHERE id = $4;
    `
	res, err := tx.ExecContext(ctx, query, item.Name, item.Description, item.Price, itemID)
	if err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to update item: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to commit tx: %w", err)
	}

	return nil
}
