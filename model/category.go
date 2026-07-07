package model

import "time"

type Category struct {
	ID        int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	OwnerId   uint64    `json:"owner_id" gorm:"column:owner_id;not null"`
	User      *User     `json:"user" gorm:"foreignKey:OwnerId;references:ID"`
}

func (c *Category) TableName() string {
	return "categories"
}
