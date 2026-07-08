package usersvc

import (
	"context"
	"errors"
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/repository/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	UserRepository userrepo.Repository
	DB             *gorm.DB
	JWTSecret      string
}

func NewServiceImpl(userRepository userrepo.Repository, db *gorm.DB, jwtSecret string) Service {
	return &ServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		JWTSecret:      jwtSecret,
	}
}

func (s *ServiceImpl) Register(ctx context.Context, user model.User) model.User {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashed)
	return s.UserRepository.CreateUser(ctx, s.DB, user)
}

func (s *ServiceImpl) Login(ctx context.Context, email string, password string) (model.User, string) {
	var user model.User
	err := s.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		panic(errors.New("invalid email or password"))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic(errors.New("invalid email or password"))
	}
	token, err := helper.GenerateToken(user.Id, user.Email, s.JWTSecret)
	helper.PanicIfError(err)
	return user, token
}

func (s *ServiceImpl) FindById(ctx context.Context, userId int) model.User {
	return s.UserRepository.FindById(ctx, s.DB, userId)
}

func (s *ServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.User, int64) {
	return s.UserRepository.FindAllUser(ctx, s.DB, page, pageSize)
}

func (s *ServiceImpl) Update(ctx context.Context, user model.User) model.User {
	return s.UserRepository.UpdateUser(ctx, s.DB, user)
}

func (s *ServiceImpl) Delete(ctx context.Context, user model.User) model.User {
	return s.UserRepository.DeleteUser(ctx, s.DB, user)
}
