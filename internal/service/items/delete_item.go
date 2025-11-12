package items

import (
	"context"
	"fmt"
)

func (s *Service) DeleteItem(ctx context.Context, itemID uint, userID uint) error {
	err := s.repo.DeleteItem(ctx, itemID, userID)
	if err != nil {
		return fmt.Errorf("service/delete_item.go - %w", err)
	}

	return nil
}
