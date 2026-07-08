package userrepo

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

func (r *RepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) model.User {
	var user model.User
	err := tx.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	helper.HandleDBError(err, "user not found")
	return user
}

func (r *RepositoryImpl) CreateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	err := tx.WithContext(ctx).Create(&user).Error
	helper.HandleDBError(err, "")
	return user
}

func (r *RepositoryImpl) UpdateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	err := tx.WithContext(ctx).Where("id = ?", user.Id).Updates(&user).Error
	helper.HandleDBError(err, "")
	return user
}

func (r *RepositoryImpl) DeleteUser(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	err := tx.WithContext(ctx).Where("id = ?", user.Id).Delete(&user).Error
	helper.HandleDBError(err, "")
	return user
}

func (r *RepositoryImpl) FindAllUser(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.User, int64) {
	var users []model.User
	var total int64
	tx.WithContext(ctx).Model(&model.User{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Limit(pageSize).Offset(offset).Find(&users).Error
	helper.HandleDBError(err, "")
	return users, total
}
