package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

func Migrate(dst ...interface{}) {
	err := DB.AutoMigrate(dst...)
	if err != nil {
		panic(err)
	}

	err = DB.Migrator().DropTable(dst...)
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(dst...)
	if err != nil {
		panic(err)
	}
}
