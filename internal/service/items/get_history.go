package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (s *Service) GetHistory(ctx context.Context, filter *models.HistoryFilter, userID uint) ([]*models.HistoryItem, error) {
	history, err := s.repo.GetHistory(ctx, filter, userID)
	if err != nil {
		return nil, fmt.Errorf("service/get_history.go: %w", err)
	}
	return history, nil
}
