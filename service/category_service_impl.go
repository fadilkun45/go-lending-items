package service

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository"

	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, db *gorm.DB) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Create(ctx, s.DB, category)
}

func (s *CategoryServiceImpl) Update(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Update(ctx, s.DB, category)
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, category model.Category) model.Category {
	return s.CategoryRepository.Delete(ctx, s.DB, category)
}

func (s *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) model.Category {
	return s.CategoryRepository.FindById(ctx, s.DB, categoryId)
}

func (s *CategoryServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Category, int64) {
	return s.CategoryRepository.FindAll(ctx, s.DB, page, pageSize)
}
