package storage

import (
	"context"
	"demo_project/modules/items/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {
	if err := s.db.Where(cond).Updates(&dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
