package config

import (
	"loans-item-go/helper"
	"loans-item-go/model"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Item{},
		&model.Loan{},
		&model.LoanItem{},
	)
	helper.PanicIfError(err)
}
