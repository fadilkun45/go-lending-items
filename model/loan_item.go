package model

import "time"

type LoanItem struct {
	Id         int64      `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	ItemId     int64      `json:"item_id" gorm:"column:item_id;not null"`
	LoanId     int64      `json:"loan_id" gorm:"column:loan_id;not null"`
	ReturnedAt *time.Time `json:"returned_at" gorm:"column:returned_at"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Item       *Item      `json:"item" gorm:"foreignKey:ItemId;references:ID"`
	Loan       *Loan      `json:"loan" gorm:"foreignKey:LoanId;references:ID"`
}

func (l *LoanItem) TableName() string {
	return "loan_items"
}
