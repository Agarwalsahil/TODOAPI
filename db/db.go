package db

import (
	"log"

	"github.com/Agarwalsahil/TodoAPI/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to databse: ", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Todo{})

	if err != nil {
		log.Fatal("Failed to migrate the tables: ", err)
	}

	log.Println("Database connection and migration successful")
}
