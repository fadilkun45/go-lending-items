package loanitemrepo

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Create(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem {
	err := tx.WithContext(ctx).Create(&loanItem).Error
	helper.HandleDBError(err, "")
	return loanItem
}

func (r *RepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, loanItemId int) model.LoanItem {
	var loanItem model.LoanItem
	err := tx.WithContext(ctx).Joins("Item").Joins("Loan").Where("loan_items.id = ?", loanItemId).First(&loanItem).Error
	helper.HandleDBError(err, "loan item not found")
	return loanItem
}

func (r *RepositoryImpl) FindReturnedLoanIDs(ctx context.Context, tx *gorm.DB, loanIDs []int64) map[int64]bool {
	var rows []model.LoanItem
	tx.WithContext(ctx).Select("loan_id").Where("loan_id IN ? AND action = ?", loanIDs, model.ActionReturn).Find(&rows)
	returned := make(map[int64]bool, len(rows))
	for _, row := range rows {
		returned[row.LoanId] = true
	}
	return returned
}

func (r *RepositoryImpl) FindByLoanID(ctx context.Context, tx *gorm.DB, loanID int, page int, pageSize int) ([]model.LoanItem, int64) {
	var loanItems []model.LoanItem
	var total int64
	tx.WithContext(ctx).Model(&model.LoanItem{}).Where("loan_id = ?", loanID).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Item").Joins("Loan").Where("loan_items.loan_id = ?", loanID).Limit(pageSize).Offset(offset).Find(&loanItems).Error
	helper.HandleDBError(err, "")
	return loanItems, total
}

func (r *RepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.LoanItem, int64) {
	var loanItems []model.LoanItem
	var total int64
	tx.WithContext(ctx).Model(&model.LoanItem{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).Joins("Item").Joins("Loan").Limit(pageSize).Offset(offset).Find(&loanItems).Error
	helper.HandleDBError(err, "")
	return loanItems, total
}
