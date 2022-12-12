package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=bookstore sslmode=disable password=jarvis22121")
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Book{})
	fmt.Println("Database connected")
	DB = db

}
