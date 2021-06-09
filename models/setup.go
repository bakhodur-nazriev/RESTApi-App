package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "user=bakhodur password=1996 dbname=test sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{}, &User{}, &Author{})

	DB = db
}
