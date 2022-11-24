package database

import (
	"fmt"
	"log"

	"github.com/shubham/bookstore/pkg/config"
	"github.com/shubham/bookstore/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := config.Config("POSTGRES_URL")

	if err != nil {
		log.Println("Idiot")
	}

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	DB.AutoMigrate(&models.Book{})
	fmt.Println("Database Migrated")

}
