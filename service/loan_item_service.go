package service

import (
	"context"
	"loans-item-go/model"
)

type LoanItemService interface {
	Create(ctx context.Context, loanItem model.LoanItem) model.LoanItem
	Update(ctx context.Context, loanItem model.LoanItem) model.LoanItem
	Delete(ctx context.Context, loanItem model.LoanItem) model.LoanItem
	FindById(ctx context.Context, loanItemId int) model.LoanItem
	FindByLoanID(ctx context.Context, loanID int, page int, pageSize int) ([]model.LoanItem, int64)
	FindAll(ctx context.Context, page int, pageSize int) ([]model.LoanItem, int64)
}
