package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (s *Service) CreateItem(ctx context.Context, item *models.ItemDTO) (uint, error) {
	id, err := s.repo.CreateItem(ctx, item)
	if err != nil {
		return 0, fmt.Errorf("service/create_item.go - %w", err)
	}

	return id, nil
}
