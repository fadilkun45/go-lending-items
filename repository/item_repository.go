package repository

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	Update(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	Delete(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	FindById(ctx context.Context, tx *gorm.DB, itemId int) model.Item
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Item, int64)
}
