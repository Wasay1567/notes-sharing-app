package controllers

import (
	"log"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err = db.AutoMigrate(&models.User{}, &models.Notes{}); err != nil {
		log.Fatal("Migration Failed ", err)
	}
	log.Println("Database Migration Successful....")
}
