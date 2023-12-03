package model

import "time"

type Stream struct {
	Id          int64     `gorm:"primaryKey"`
	Description string    `gorm:"description;not null;"`
	Code        string    `gorm:"code;unique;not null;"`
	CreatedDt   time.Time `gorm:"created_dt,not null;default:now();"`
}

func (u *Stream) TableName() string {
	return "user.stream"
}
