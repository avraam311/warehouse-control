package items

import (
	"context"
	"fmt"

	"github.com/avraam311/warehouse-control/internal/models"
)

func (s *Service) ReplaceItem(ctx context.Context, itemID uint, item *models.ItemDTO, userID uint) error {
	err := s.repo.ReplaceItem(ctx, itemID, item, userID)
	if err != nil {
		return fmt.Errorf("service/replace_item.go - %w", err)
	}

	return nil
}
