package categorysvc

import (
	"context"
	"loans-item-go/model"
)

type Service interface {
	Create(ctx context.Context, category model.Category) model.Category
	Update(ctx context.Context, category model.Category) model.Category
	Delete(ctx context.Context, category model.Category) model.Category
	FindById(ctx context.Context, categoryId int) model.Category
	FindAll(ctx context.Context, page int, pageSize int) ([]model.Category, int64)
	Search(ctx context.Context, query string, page int, pageSize int) ([]model.Category, int64)
}
