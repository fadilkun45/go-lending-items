package repository

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepositoryImpl() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (r *CategoryRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Create(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Where("id = ?", category.ID).Updates(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Where("id = ?", category.ID).Delete(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *CategoryRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) model.Category {
	var category model.Category
	err := tx.WithContext(ctx).Joins("User").Where("categories.id = ?", categoryId).First(&category).Error
	helper.HandleDBError(err, "category not found")
	return category
}

func (r *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Category, int64) {
	var categories []model.Category
	var total int64
	tx.WithContext(ctx).Model(&model.Category{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("User").Limit(pageSize).Offset(offset).Find(&categories).Error
	helper.HandleDBError(err, "")
	return categories, total
}
