package models

import "time"

type Photo struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string    `gorm:"column:title;type:varchar"`
	Caption   string    `gorm:"column:caption;type:varchar"`
	PhotoUrl  string    `gorm:"column:photo_url;type:varchar"`
	UserID    int       `gorm:"column:user_id;type:int"`
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User      *User
}
