package categoryrepo

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Create(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Create(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Where("id = ?", category.ID).Updates(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category model.Category) model.Category {
	err := tx.WithContext(ctx).Where("id = ?", category.ID).Delete(&category).Error
	helper.HandleDBError(err, "")
	return category
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, categoryId int) model.Category {
	var category model.Category
	err := tx.WithContext(ctx).Where("id = ?", categoryId).First(&category).Error
	helper.HandleDBError(err, "category not found")
	return category
}

func (r *RepositoryImpl) Search(ctx context.Context, tx *gorm.DB, query string, page int, pageSize int) ([]model.Category, int64) {
	var categories []model.Category
	var total int64
	pattern := "%" + query + "%"
	tx.WithContext(ctx).Model(&model.Category{}).Where("name LIKE ?", pattern).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Where("name LIKE ?", pattern).Limit(pageSize).Offset(offset).Find(&categories).Error
	helper.HandleDBError(err, "")
	return categories, total
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Category, int64) {
	var categories []model.Category
	var total int64
	tx.WithContext(ctx).Model(&model.Category{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Owner").Limit(pageSize).Offset(offset).Find(&categories).Error
	helper.HandleDBError(err, "")
	return categories, total
}
