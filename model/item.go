package model

import "time"

type Item struct {
	ID         int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name       string    `json:"name" gorm:"column:name;not null"`
	CategoryID int64     `json:"category_id" gorm:"column:category_id;not null"`
	Category   *Category `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	SerialNo   *string   `json:"serial_no" gorm:"column:serial_no;type:varchar(255);uniqueIndex"`
	Condition  string    `json:"condition" gorm:"column:condition;"`
	Status     string    `json:"status" gorm:"column:status;default:available"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	OwnerID    int64     `json:"owner_id" gorm:"column:owner_id;not null"`
	Owner      *User     `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}

func (i *Item) TableName() string {
	return "items"
}
