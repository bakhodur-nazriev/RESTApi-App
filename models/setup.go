package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Books{}, &Users{})

	DB = db
}
