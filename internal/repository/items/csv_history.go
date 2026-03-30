package items

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (r *Repository) ExportHistoryCSV(ctx context.Context, filter *models.HistoryFilter, userID uint) ([]byte, error) {
	history, err := r.GetHistory(ctx, filter, userID)
	if err != nil {
		return nil, fmt.Errorf("csv_history.go - failed to get history: %w", err)
	}

	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	if err := writer.Write([]string{"history_id", "item_id", "action", "changed_at", "changed_by_email", "name", "description", "price"}); err != nil {
		return nil, fmt.Errorf("csv_history.go - failed to write header: %w", err)
	}

	for _, h := range history {
		row := []string{
			fmt.Sprintf("%d", h.HistoryID),
			fmt.Sprintf("%d", h.ItemID),
			h.Action,
			h.ChangedAt.Format("2006-01-02 15:04:05"),
			h.ChangedByEmail,
			h.Name,
			strings.ReplaceAll(h.Description, "\"", "\"\""), 
			fmt.Sprintf("%.2f", h.Price),
		}
		if err := writer.Write(row); err != nil {
			return nil, fmt.Errorf("csv_history.go - failed to write row: %w", err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, fmt.Errorf("csv_history.go - writer flush error: %w", err)
	}

	return buf.Bytes(), nil
}
