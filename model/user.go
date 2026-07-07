package model

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey;autoIncrement:true"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	Email     string    `json:"email" gorm:"column:email;not null;unique"`
	Password  string    `json:"-" gorm:"column:password;not null" `
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}
