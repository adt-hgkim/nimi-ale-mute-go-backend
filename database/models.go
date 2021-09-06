package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       uint32 `gorm:"primaryKey;autoIncrement:true"`
	Email    string `gorm:"unique"`
	Password string
	Access   string
	Refresh  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Word struct {
	Id uint64 `gorm:"primaryKey;autoIncrement:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
