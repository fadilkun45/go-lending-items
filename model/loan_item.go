package model

import "time"

const (
	ActionBorrow = "BORROW"
	ActionReturn = "RETURN"
)

type LoanItem struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	ItemId    int64     `json:"item_id" gorm:"column:item_id;not null"`
	LoanId    int64     `json:"loan_id" gorm:"column:loan_id;not null"`
	Action    string    `json:"action" gorm:"column:action;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Item      *Item     `json:"item" gorm:"foreignKey:ItemId;references:ID"`
	Loan      *Loan     `json:"loan" gorm:"foreignKey:LoanId;references:ID"`
}

func (l *LoanItem) TableName() string {
	return "loan_items"
}
