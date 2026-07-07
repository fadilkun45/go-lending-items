package service

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository"

	"gorm.io/gorm"
)

type ItemServiceImpl struct {
	ItemRepository repository.ItemRepository
	DB             *gorm.DB
}

func NewItemServiceImpl(itemRepository repository.ItemRepository, db *gorm.DB) ItemService {
	return &ItemServiceImpl{
		ItemRepository: itemRepository,
		DB:             db,
	}
}

func (s *ItemServiceImpl) Create(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Create(ctx, s.DB, item)
}

func (s *ItemServiceImpl) Update(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Update(ctx, s.DB, item)
}

func (s *ItemServiceImpl) Delete(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Delete(ctx, s.DB, item)
}

func (s *ItemServiceImpl) FindById(ctx context.Context, itemId int) model.Item {
	return s.ItemRepository.FindById(ctx, s.DB, itemId)
}

func (s *ItemServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Item, int64) {
	return s.ItemRepository.FindAll(ctx, s.DB, page, pageSize)
}
