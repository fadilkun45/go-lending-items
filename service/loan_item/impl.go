package loanitemsvc

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository/loan_item"

	"gorm.io/gorm"
)

type ServiceImpl struct {
	LoanItemRepository loanitemrepo.Repository
	DB                 *gorm.DB
}

func NewServiceImpl(loanItemRepository loanitemrepo.Repository, db *gorm.DB) Service {
	return &ServiceImpl{
		LoanItemRepository: loanItemRepository,
		DB:                 db,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, loanItem model.LoanItem) model.LoanItem {
	return s.LoanItemRepository.Create(ctx, s.DB, loanItem)
}

func (s *ServiceImpl) FindById(ctx context.Context, loanItemId int) model.LoanItem {
	return s.LoanItemRepository.FindById(ctx, s.DB, loanItemId)
}

func (s *ServiceImpl) FindByLoanID(ctx context.Context, loanID int, page int, pageSize int) ([]model.LoanItem, int64) {
	return s.LoanItemRepository.FindByLoanID(ctx, s.DB, loanID, page, pageSize)
}

func (s *ServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.LoanItem, int64) {
	return s.LoanItemRepository.FindAll(ctx, s.DB, page, pageSize)
}
