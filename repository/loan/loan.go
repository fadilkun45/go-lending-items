package loanrepo

import (
	"context"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan
	Update(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan
	Delete(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan
	FindById(ctx context.Context, tx *gorm.DB, loanId int) model.Loan
	FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Loan, int64)
	FindByBorrowerID(ctx context.Context, tx *gorm.DB, borrowerID int, page int, pageSize int) ([]model.Loan, int64)
}
