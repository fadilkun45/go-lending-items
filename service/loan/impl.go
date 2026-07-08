package loansvc

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"
	itemrepo "loans-item-go/repository/item"
	loanrepo "loans-item-go/repository/loan"
	loanitemrepo "loans-item-go/repository/loan_item"

	"gorm.io/gorm"
)

type ServiceImpl struct {
	LoanRepository     loanrepo.Repository
	ItemRepository     itemrepo.Repository
	LoanItemRepository loanitemrepo.Repository
	DB                 *gorm.DB
}

func NewServiceImpl(loanRepository loanrepo.Repository, itemRepository itemrepo.Repository, loanItemRepository loanitemrepo.Repository, db *gorm.DB) Service {
	return &ServiceImpl{
		LoanRepository:     loanRepository,
		ItemRepository:     itemRepository,
		LoanItemRepository: loanItemRepository,
		DB:                 db,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, loan model.Loan) model.Loan {
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
			Action: model.ActionBorrow,
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

func (s *ServiceImpl) Update(ctx context.Context, loan model.Loan) model.Loan {
	var result model.Loan
	err := s.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		existing := s.LoanRepository.FindById(ctx, tx, int(loan.ID))
		result = s.LoanRepository.Update(ctx, tx, loan)
		s.LoanItemRepository.Create(ctx, tx, model.LoanItem{
			ItemId: existing.ItemID,
			LoanId: loan.ID,
			Action: model.ActionReturn,
		})
		s.ItemRepository.Update(ctx, tx, model.Item{
			ID:     existing.ItemID,
			Status: "AVAILABLE",
		})
		return nil
	})
	helper.HandleDBError(err, "")
	return result
}

func (s *ServiceImpl) Delete(ctx context.Context, loan model.Loan) model.Loan {
	return s.LoanRepository.Delete(ctx, s.DB, loan)
}

func (s *ServiceImpl) FindById(ctx context.Context, loanId int) model.Loan {
	loans := []model.Loan{s.LoanRepository.FindById(ctx, s.DB, loanId)}
	s.attachStatuses(ctx, loans)
	return loans[0]
}

func (s *ServiceImpl) FindAll(ctx context.Context, page int, pageSize int) ([]model.Loan, int64) {
	loans, total := s.LoanRepository.FindAll(ctx, s.DB, page, pageSize)
	s.attachStatuses(ctx, loans)
	return loans, total
}

func (s *ServiceImpl) FindByBorrowerID(ctx context.Context, borrowerID int, page int, pageSize int) ([]model.Loan, int64) {
	loans, total := s.LoanRepository.FindByBorrowerID(ctx, s.DB, borrowerID, page, pageSize)
	s.attachStatuses(ctx, loans)
	return loans, total
}

// attachStatuses derives each loan's status from its own loan_items history
// (the append-only borrow/return log), instead of the item's current status,
// since the item's status is shared/overwritten across every loan of that item.
func (s *ServiceImpl) attachStatuses(ctx context.Context, loans []model.Loan) {
	ids := make([]int64, len(loans))
	for i, l := range loans {
		ids[i] = l.ID
	}
	returned := s.LoanItemRepository.FindReturnedLoanIDs(ctx, s.DB, ids)
	for i := range loans {
		if returned[loans[i].ID] {
			loans[i].Status = "RETURNED"
		} else {
			loans[i].Status = "BORROWED"
		}
	}
}
