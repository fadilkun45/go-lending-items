package itemsvc

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository/item"

	"gorm.io/gorm"
)

type ServiceImpl struct {
	ItemRepository itemrepo.Repository
	DB             *gorm.DB
}

func NewServiceImpl(itemRepository itemrepo.Repository, db *gorm.DB) Service {
	return &ServiceImpl{
		ItemRepository: itemRepository,
		DB:             db,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Create(ctx, s.DB, item)
}

func (s *ServiceImpl) Update(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Update(ctx, s.DB, item)
}

func (s *ServiceImpl) Delete(ctx context.Context, item model.Item) model.Item {
	return s.ItemRepository.Delete(ctx, s.DB, item)
}

func (s *ServiceImpl) FindById(ctx context.Context, itemId int) model.Item {
	return s.ItemRepository.FindById(ctx, s.DB, itemId)
}

func (s *ServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Item, int64) {
	return s.ItemRepository.FindAll(ctx, s.DB, page, pageSize)
}

func (s *ServiceImpl) FindByOwner(ctx context.Context, ownerId int, page int, pageSize int) ([]model.Item, int64) {
	return s.ItemRepository.FindByOwner(ctx, s.DB, ownerId, page, pageSize)
}

func (s *ServiceImpl) Search(ctx context.Context, query string, categoryId int, page int, pageSize int) ([]model.Item, int64) {
	return s.ItemRepository.Search(ctx, s.DB, query, categoryId, page, pageSize)
}
