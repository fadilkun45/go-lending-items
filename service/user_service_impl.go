package service

import (
	"context"
	"errors"
	"loans-item-go/model"
	"loans-item-go/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserServiceImpl(userRepository repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, user model.User) model.User {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashed)
	return s.UserRepository.CreateUser(ctx, s.DB, user)
}

func (s *UserServiceImpl) Login(ctx context.Context, email string, password string) model.User {
	var user model.User
	err := s.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		panic(errors.New("invalid email or password"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic(errors.New("invalid email or password"))
	}

	return user
}

func (s *UserServiceImpl) FindById(ctx context.Context, userId int) model.User {
	return s.UserRepository.FindById(ctx, s.DB, userId)
}

func (s *UserServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.User, int64) {
	return s.UserRepository.FindAllUser(ctx, s.DB, page, pageSize)
}

func (s *UserServiceImpl) Update(ctx context.Context, user model.User) model.User {
	return s.UserRepository.UpdateUser(ctx, s.DB, user)
}

func (s *UserServiceImpl) Delete(ctx context.Context, user model.User) model.User {
	return s.UserRepository.DeleteUser(ctx, s.DB, user)
}
