package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TabelName() string {
	return "user"
}
