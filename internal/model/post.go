package model

import "time"

type Post struct {
	Id          int64     `gorm:"primaryKey"`
	Description string    `gorm:"description;not null;"`
	Code        string    `gorm:"code;unique;not null;"`
	CreatedDt   time.Time `gorm:"created_dt,not null;default:now();"`
}

func (u *Post) TableName() string {
	return "user.post"
}
