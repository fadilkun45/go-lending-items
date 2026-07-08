package loanitemrepo

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem
	FindById(ctx context.Context, tx *gorm.DB, loanItemId int) model.LoanItem
	FindReturnedLoanIDs(ctx context.Context, tx *gorm.DB, loanIDs []int64) map[int64]bool
	FindByLoanID(ctx context.Context, tx *gorm.DB, loanID int, page int, pageSize int) ([]model.LoanItem, int64)
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.LoanItem, int64)
}
