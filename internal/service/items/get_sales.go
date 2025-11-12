package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (s *Service) GetItems(ctx context.Context, userID uint) ([]*models.ItemDB, error) {
	items, err := s.repo.GetItems(ctx, userID)
	if err != nil {
		return items, fmt.Errorf("service/get_items.go - %w", err)
	}

	return items, nil
}
