package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
	"github.com/wb-go/wbf/zlog"
)

func (r *Repository) GetHistory(ctx context.Context, filter *models.HistoryFilter, userID uint) ([]*models.HistoryItem, error) {
	tx, err := r.db.Master.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("get_history.go - failed to begin tx: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	setUserIDQuery := fmt.Sprintf("SET LOCAL myapp.current_user_id = %d", userID)
	if _, err = tx.ExecContext(ctx, setUserIDQuery); err != nil {
		return nil, fmt.Errorf("get_history.go - failed to set local user_id: %w", err)
	}

	query := `
		SELECT 
			ih.history_id, ih.item_id, ih.action, ih.changed_at, ih.changed_by, 
			u.email, ih.name, ih.description, ih.price
		FROM item_history ih
		LEFT JOIN "user" u ON ih.changed_by = u.id
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	if filter.ItemID != nil {
		query += fmt.Sprintf(" AND ih.item_id = $%d", argIndex)
		args = append(args, *filter.ItemID)
		argIndex++
	}
	if filter.UserID != nil {
		query += fmt.Sprintf(" AND ih.changed_by = $%d", argIndex)
		args = append(args, *filter.UserID)
		argIndex++
	}
	if filter.Action != nil {
		query += fmt.Sprintf(" AND ih.action = $%d", argIndex)
		args = append(args, *filter.Action)
		argIndex++
	}
	if filter.DateFrom != nil {
		query += fmt.Sprintf(" AND ih.changed_at >= $%d", argIndex)
		args = append(args, *filter.DateFrom)
		argIndex++
	}
	if filter.DateTo != nil {
		query += fmt.Sprintf(" AND ih.changed_at <= $%d", argIndex)
		args = append(args, *filter.DateTo)
		argIndex++
	}

	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, *filter.Limit)
		argIndex++
	}
	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, *filter.Offset)
	}

	query += " ORDER BY ih.changed_at DESC"

	zlog.Logger.Info().Str("query", query).Msg("executing history query")

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_history.go - query failed: %w", err)
	}
	defer rows.Close()

	var history []*models.HistoryItem
	for rows.Next() {
		var h models.HistoryItem
		err := rows.Scan(
			&h.HistoryID, &h.ItemID, &h.Action, &h.ChangedAt,
			&h.ChangedBy, &h.ChangedByEmail, &h.Name, &h.Description, &h.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("get_history.go - scan failed: %w", err)
		}
		history = append(history, &h)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("get_history.go - rows error: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("get_history.go - commit failed: %w", err)
	}

	return history, nil
}
