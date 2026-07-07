package service

import (
	"context"
	"loans-item-go/model"
)

type LoanService interface {
	Create(ctx context.Context, loan model.Loan) model.Loan
	Update(ctx context.Context, loan model.Loan) model.Loan
	Delete(ctx context.Context, loan model.Loan) model.Loan
	FindById(ctx context.Context, loanId int) model.Loan
	FindAll(ctx context.Context, page int, pageSize int) ([]model.Loan, int64)
	FindByBorrowerID(ctx context.Context, borrowerID int, page int, pageSize int) ([]model.Loan, int64)
}
