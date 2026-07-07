package repository

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(ctx context.Context, tx *gorm.DB, userId int) model.User
	CreateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User
	UpdateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User
	DeleteUser(ctx context.Context, tx *gorm.DB, user model.User) model.User
	FindAllUser(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.User, int64)
}
