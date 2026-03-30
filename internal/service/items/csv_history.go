package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (s *Service) ExportHistoryCSV(ctx context.Context, filter *models.HistoryFilter, userID uint) ([]byte, error) {
	data, err := s.repo.ExportHistoryCSV(ctx, filter, userID)
	if err != nil {
		return nil, fmt.Errorf("service/csv_history.go: %w", err)
	}
	return data, nil
}
