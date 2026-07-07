package service

import (
	"context"
	"loans-item-go/model"
)

type ItemService interface {
	Create(ctx context.Context, item model.Item) model.Item
	Update(ctx context.Context, item model.Item) model.Item
	Delete(ctx context.Context, item model.Item) model.Item
	FindById(ctx context.Context, itemId int) model.Item
	FindAll(ctx context.Context, page int, pageSize int) ([]model.Item, int64)
}
