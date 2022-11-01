package models

import "time"

type User struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `gorm:"column:username;type:varchar;size:255;unique;not null"`
	Email     string `gorm:"column:email;type:varchar;size:255;unique;not null"`
	Password  string `gorm:"column:password;not null"`
	Age       int    `gorm:"column:age;type:int;not null"`
}
