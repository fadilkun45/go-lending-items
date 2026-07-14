package categorysvc

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository/category"

	"gorm.io/gorm"
)

type ServiceImpl struct {
	CategoryRepository categoryrepo.Repository
	DB                 *gorm.DB
}

func NewServiceImpl(categoryRepository categoryrepo.Repository, db *gorm.DB) Service {
	return &ServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Create(ctx, s.DB, category)
}

func (s *ServiceImpl) Update(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Update(ctx, s.DB, category)
}

func (s *ServiceImpl) Delete(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Delete(ctx, s.DB, category)
}

func (s *ServiceImpl) FindById(ctx context.Context, categoryId int) model.Category {
	return s.CategoryRepository.FindById(ctx, s.DB, categoryId)
}

func (s *ServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Category, int64) {
	return s.CategoryRepository.FindAll(ctx, s.DB, page, pageSize)
}

func (s *ServiceImpl) Search(ctx context.Context, query string, page int, pageSize int) ([]model.Category, int64) {
	return s.CategoryRepository.Search(ctx, s.DB, query, page, pageSize)
}
