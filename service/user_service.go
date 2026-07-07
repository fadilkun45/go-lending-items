package service

import (
	"context"
	"loans-item-go/model"
)

type UserService interface {
	Register(ctx context.Context, user model.User) model.User
	Login(ctx context.Context, email string, password string) model.User
	FindById(ctx context.Context, userId int) model.User
	FindAll(ctx context.Context, page int, pageSize int) ([]model.User, int64)
	Update(ctx context.Context, user model.User) model.User
	Delete(ctx context.Context, user model.User) model.User
}
