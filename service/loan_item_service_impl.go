package service

import (
	"context"
	"loans-item-go/model"
	"loans-item-go/repository"

	"gorm.io/gorm"
)

type LoanItemServiceImpl struct {
	LoanItemRepository repository.LoanItemRepository
	DB                 *gorm.DB
}

func NewLoanItemServiceImpl(loanItemRepository repository.LoanItemRepository, db *gorm.DB) LoanItemService {
	return &LoanItemServiceImpl{
		LoanItemRepository: loanItemRepository,
		DB:                 db,
	}
}

func (s *LoanItemServiceImpl) Create(ctx context.Context, loanItem model.LoanItem) model.LoanItem {
	return s.LoanItemRepository.Create(ctx, s.DB, loanItem)
}

func (s *LoanItemServiceImpl) Update(ctx context.Context, loanItem model.LoanItem) model.LoanItem {
	return s.LoanItemRepository.Update(ctx, s.DB, loanItem)
}

func (s *LoanItemServiceImpl) Delete(ctx context.Context, loanItem model.LoanItem) model.LoanItem {
	return s.LoanItemRepository.Delete(ctx, s.DB, loanItem)
}

func (s *LoanItemServiceImpl) FindById(ctx context.Context, loanItemId int) model.LoanItem {
	return s.LoanItemRepository.FindById(ctx, s.DB, loanItemId)
}

func (s *LoanItemServiceImpl) FindByLoanID(ctx context.Context, loanID int, page int, pageSize int) ([]model.LoanItem, int64) {
	return s.LoanItemRepository.FindByLoanID(ctx, s.DB, loanID, page, pageSize)
}

func (s *LoanItemServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.LoanItem, int64) {
	return s.LoanItemRepository.FindAll(ctx, s.DB, page, pageSize)
}
