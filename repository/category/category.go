package categoryrepo

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindById(ctx context.Context, tx *gorm.DB, categoryId int) model.Category
	Create(ctx context.Context, tx *gorm.DB, category model.Category) model.Category
	Update(ctx context.Context, tx *gorm.DB, category model.Category) model.Category
	Delete(ctx context.Context, tx *gorm.DB, category model.Category) model.Category
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Category, int64)
	Search(ctx context.Context, tx *gorm.DB, query string, page int, pageSize int) ([]model.Category, int64)
}
