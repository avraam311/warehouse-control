package items

import (
	"context"
	"fmt"
)

func (s *Service) DeleteItem(ctx context.Context, id uint) error {
	err := s.repo.DeleteItem(ctx, id)
	if err != nil {
		return fmt.Errorf("service/delete_item.go - %w", err)
	}

	return nil
}
