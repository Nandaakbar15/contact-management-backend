package models

// package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/contact_db"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Contact{})

	DB = database
}
