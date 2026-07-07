package service

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"
	"loans-item-go/repository"
	"time"

	"gorm.io/gorm"
)

type LoanServiceImpl struct {
	LoanRepository     repository.LoanRepository
	ItemRepository     repository.ItemRepository
	LoanItemRepository repository.LoanItemRepository
	DB                 *gorm.DB
}

func NewLoanServiceImpl(loanRepository repository.LoanRepository, itemRepository repository.ItemRepository, loanItemRepository repository.LoanItemRepository, db *gorm.DB) LoanService {
	return &LoanServiceImpl{
		LoanRepository:     loanRepository,
		ItemRepository:     itemRepository,
		LoanItemRepository: loanItemRepository,
		DB:                 db,
	}
}

func (s *LoanServiceImpl) Create(ctx context.Context, loan model.Loan) model.Loan {
	var result model.Loan
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		item := s.ItemRepository.FindById(ctx, tx, int(loan.ItemID))
		if item.Status != "AVAILABLE" {
			panic(helper.BadRequest("item is not available"))
		}
		result = s.LoanRepository.Create(ctx, tx, loan)
		s.LoanItemRepository.Create(ctx, tx, model.LoanItem{
			ItemId: loan.ItemID,
			LoanId: result.ID,
		})
		s.ItemRepository.Update(ctx, tx, model.Item{
			ID:     loan.ItemID,
			Status: "BORROWED",
		})
		return nil
	})
	helper.HandleDBError(err, "")
	return result
}

func (s *LoanServiceImpl) Update(ctx context.Context, loan model.Loan) model.Loan {
	var result model.Loan
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		existing := s.LoanRepository.FindById(ctx, tx, int(loan.ID))
		result = s.LoanRepository.Update(ctx, tx, loan)
		returnedAt := time.Now()
		if loan.ReturnedAt != nil {
			returnedAt = *loan.ReturnedAt
		}
		s.LoanItemRepository.ReturnByLoanID(ctx, tx, loan.ID, returnedAt)
		s.ItemRepository.Create(ctx, tx, model.Item{
			ID:     existing.ItemID,
			Status: "AVAILABLE",
		})
		return nil
	})
	helper.HandleDBError(err, "")
	return result
}

func (s *LoanServiceImpl) Delete(ctx context.Context, loan model.Loan) model.Loan {
	return s.LoanRepository.Delete(ctx, s.DB, loan)
}

func (s *LoanServiceImpl) FindById(ctx context.Context, loanId int) model.Loan {
	return s.LoanRepository.FindById(ctx, s.DB, loanId)
}

func (s *LoanServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Loan, int64) {
	return s.LoanRepository.FindAll(ctx, s.DB, page, pageSize)
}

func (s *LoanServiceImpl) FindByBorrowerID(ctx context.Context, borrowerID int, page int, pageSize int) ([]model.Loan, int64) {
	return s.LoanRepository.FindByBorrowerID(ctx, s.DB, borrowerID, page, pageSize)
}
