package repository

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type ItemRepositoryImpl struct{}

func NewItemRepositoryImpl() ItemRepository {
	return &ItemRepositoryImpl{}
}

func (r *ItemRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Create(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *ItemRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Where("id = ?", item.ID).Updates(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *ItemRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Where("id = ?", item.ID).Delete(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *ItemRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, itemId int) model.Item {
	var item model.Item
	err := tx.WithContext(ctx).Joins("Category").Joins("Owner").Where("items.id = ?", itemId).First(&item).Error
	helper.HandleDBError(err, "item not found")
	return item
}

func (r *ItemRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Item, int64) {
	var items []model.Item
	var total int64
	tx.WithContext(ctx).Model(&model.Item{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Category.User").Joins("Owner").Limit(pageSize).Offset(offset).Find(&items).Error
	helper.HandleDBError(err, "")
	return items, total
}
