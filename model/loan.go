package model

import "time"

type Loan struct {
	ID         int64      `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	ItemID     int64      `json:"item_id" gorm:"column:item_id;not null"`
	Item       Item       `json:"item" gorm:"foreignKey:ItemID;references:ID"`
	BorrowerID int64      `json:"borrower_id" gorm:"column:borrower_id;not null"`
	Borrower   User       `json:"borrower" gorm:"foreignKey:BorrowerID;references:ID"`
	BorrowedAt time.Time  `json:"borrowed_at" gorm:"column:borrowed_at;autoCreateTime"`
	ReturnedAt *time.Time `json:"returned_at" gorm:"column:returned_at"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Status     string     `json:"status" gorm:"-"`
}

func (l *Loan) TableName() string {
	return "loans"
}
