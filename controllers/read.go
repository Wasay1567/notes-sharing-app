package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"gorm.io/gorm"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	pass := r.PathValue("pass")

	var note models.Notes
	result := db.First(&note, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		log.Printf("Database error: %v", result.Error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if note.Password != pass {
		http.Error(w, "Incorrect Password", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)

	note.Viewed = true
	db.Save(&note)
	if note.Expiration == "1 view" && note.Viewed {
		note.IsExpired = true
		db.Save(&note)
	}
}
