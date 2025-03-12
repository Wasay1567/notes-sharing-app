package controllers

import (
	"fmt"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}
	db.AutoMigrate(&models.Notes{})
	fmt.Println("Database Connected Successfully....")
}
