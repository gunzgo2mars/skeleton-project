package model

import "time"

type UserSchema struct {
	ID        uint      `gorm:"id"`
	UUID      string    `gorm:"uuid"`
	Firstname string    `gorm:"firstname"`
	Lastname  string    `gorm:"lastname"`
	CreateAt  time.Time `gorm:"create_at"`
	UpdateAt  time.Time `gorm:"update_at"`
	IsActive  bool      `gorm:"is_active"`
}
