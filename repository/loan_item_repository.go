package repository

import (
	"context"
	"loans-item-go/model"
	"time"

	"gorm.io/gorm"
)

type LoanItemRepository interface {
	Create(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem
	Update(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem
	Delete(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem
	FindById(ctx context.Context, tx *gorm.DB, loanItemId int) model.LoanItem
	ReturnByLoanID(ctx context.Context, tx *gorm.DB, loanID int64, returnedAt time.Time)
	FindByLoanID(ctx context.Context, tx *gorm.DB, loanID int, page int, pageSize int) ([]model.LoanItem, int64)
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.LoanItem, int64)
}
