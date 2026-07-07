package repository

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type LoanRepositoryImpl struct{}

func NewLoanRepositoryImpl() LoanRepository {
	return &LoanRepositoryImpl{}
}

func (r *LoanRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan {
	err := tx.WithContext(ctx).Create(&loan).Error
	helper.HandleDBError(err, "")
	return loan
}

func (r *LoanRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan {
	err := tx.WithContext(ctx).Where("id = ?", loan.ID).Updates(&loan).Error
	helper.HandleDBError(err, "")
	return loan
}

func (r *LoanRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, loan model.Loan) model.Loan {
	err := tx.WithContext(ctx).Where("id = ?", loan.ID).Delete(&loan).Error
	helper.HandleDBError(err, "")
	return loan
}

func (r *LoanRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, loanId int) model.Loan {
	var loan model.Loan
	err := tx.WithContext(ctx).Joins("Borrower").Joins("Item").Joins("Item.Owner").Joins("Item.Category").Where("loans.id = ?", loanId).First(&loan).Error
	helper.HandleDBError(err, "loan not found")
	return loan
}

func (r *LoanRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.Loan, int64) {
	var loans []model.Loan
	var total int64
	tx.WithContext(ctx).Model(&model.Loan{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Borrower").Joins("Item.Owner").Joins("Item.Category").Limit(pageSize).Offset(offset).Find(&loans).Error
	helper.HandleDBError(err, "")
	return loans, total
}

func (r *LoanRepositoryImpl) FindByBorrowerID(ctx context.Context, tx *gorm.DB, borrowerID int, page int, pageSize int) ([]model.Loan, int64) {
	var loans []model.Loan
	var total int64
	tx.WithContext(ctx).Model(&model.Loan{}).Where("borrower_id = ?", borrowerID).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Borrower").Joins("Item.Owner").Joins("Item.Category").Where("loan.borrower_id = ?", borrowerID).Limit(pageSize).Offset(offset).Find(&loans).Error
	helper.HandleDBError(err, "")
	return loans, total
}
