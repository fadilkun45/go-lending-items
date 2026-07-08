package itemrepo

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

func (r *RepositoryImpl) Create(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Create(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *RepositoryImpl) Update(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Where("id = ?", item.ID).Updates(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *RepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, item model.Item) model.Item {
	err := tx.WithContext(ctx).Where("id = ?", item.ID).Delete(&item).Error
	helper.HandleDBError(err, "")
	return item
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, itemId int) model.Item {
	var item model.Item
	err := tx.WithContext(ctx).Joins("Category").Joins("Owner").Where("items.id = ?", itemId).First(&item).Error
	helper.HandleDBError(err, "item not found")
	return item
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Item, int64) {
	var items []model.Item
	var total int64
	tx.WithContext(ctx).Model(&model.Item{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Category").Joins("Owner").Limit(pageSize).Offset(offset).Find(&items).Error
	helper.HandleDBError(err, "")
	return items, total
}

func (r *RepositoryImpl) FindByOwner(ctx context.Context, tx *gorm.DB, ownerId int, page int, pageSize int) ([]model.Item, int64) {
	var items []model.Item
	var total int64
	tx.WithContext(ctx).Model(&model.Item{}).Where("owner_id = ?", ownerId).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Category").Joins("Owner").Where("items.owner_id = ?", ownerId).Limit(pageSize).Offset(offset).Find(&items).Error
	helper.HandleDBError(err, "item not found")
	return items, total
}
