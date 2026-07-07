package repository

import (
	"context"
	"loans-item-go/helper"
	"loans-item-go/model"
	"time"

	"gorm.io/gorm"
)

type LoanItemRepositoryImpl struct{}

func NewLoanItemRepositoryImpl() LoanItemRepository {
	return &LoanItemRepositoryImpl{}
}

func (r *LoanItemRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem {
	err := tx.WithContext(ctx).Create(&loanItem).Error
	helper.HandleDBError(err, "")
	return loanItem
}

func (r *LoanItemRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem {
	err := tx.WithContext(ctx).Where("id = ?", loanItem.Id).Updates(&loanItem).Error
	helper.HandleDBError(err, "")
	return loanItem
}

func (r *LoanItemRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, loanItem model.LoanItem) model.LoanItem {
	err := tx.WithContext(ctx).Where("id = ?", loanItem.Id).Delete(&loanItem).Error
	helper.HandleDBError(err, "")
	return loanItem
}

func (r *LoanItemRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, loanItemId int) model.LoanItem {
	var loanItem model.LoanItem
	err := tx.WithContext(ctx).Joins("Item").Joins("Loan").Where("loan_item.id = ?", loanItemId).First(&loanItem).Error
	helper.HandleDBError(err, "loan item not found")
	return loanItem
}

func (r *LoanItemRepositoryImpl) ReturnByLoanID(ctx context.Context, tx *gorm.DB, loanID int64, returnedAt time.Time) {
	err := tx.WithContext(ctx).Model(&model.LoanItem{}).
		Where("loan_id = ?", loanID).
		Update("returned_at", returnedAt).Error
	helper.HandleDBError(err, "")
}

func (r *LoanItemRepositoryImpl) FindByLoanID(ctx context.Context, tx *gorm.DB, loanID int, page int, pageSize int) ([]model.LoanItem, int64) {
	var loanItems []model.LoanItem
	var total int64
	tx.WithContext(ctx).Model(&model.LoanItem{}).Where("loan_id = ?", loanID).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).
		Joins("Item").Joins("Loan").
		Where("loan_item.loan_id = ?", loanID).
		Limit(pageSize).
		Offset(offset).
		Find(&loanItems).Error
	helper.HandleDBError(err, "")
	return loanItems, total
}

func (r *LoanItemRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB, page int, pageSize int) ([]model.LoanItem, int64) {
	var loanItems []model.LoanItem
	var total int64
	tx.WithContext(ctx).Model(&model.LoanItem{}).Count(&total)
	offset := (page - 1) * pageSize
	err := tx.WithContext(ctx).
		Joins("Item").Joins("Loan").
		Limit(pageSize).
		Offset(offset).
		Find(&loanItems).Error
	helper.HandleDBError(err, "")
	return loanItems, total
}
