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

	Created time.Time `gorm:"default:null"`
	Updated time.Time `gorm:"default:null"`
	Deleted gorm.DeletedAt
}

type Word struct {
	Id uint64 `gorm:"primaryKey;autoIncrement:true"`

	Created time.Time `gorm:"default:null"`
	Updated time.Time `gorm:"default:null"`
	Deleted gorm.DeletedAt
}
