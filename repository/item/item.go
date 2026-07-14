package itemrepo

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	Update(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	Delete(ctx context.Context, tx *gorm.DB, item model.Item) model.Item
	FindById(ctx context.Context, tx *gorm.DB, itemId int) model.Item
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Item, int64)
	FindByOwner(ctx context.Context, tx *gorm.DB, ownerId int, page int, pageSize int) ([]model.Item, int64)
	Search(ctx context.Context, tx *gorm.DB, query string, categoryId int, page int, pageSize int) ([]model.Item, int64)
}
