package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) ReplaceItem(ctx context.Context, id uint, item *models.ItemDTO) error {
	query := `
        UPDATE item
        SET name = $1, description = $2, price = $3
        WHERE id = $4;
    `

	res, err := r.db.ExecContext(ctx, query, item.Name, item.Description, item.Price, id)
	if err != nil {
		return fmt.Errorf("repository/replace_item.go - failed to update item: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	return nil
}
