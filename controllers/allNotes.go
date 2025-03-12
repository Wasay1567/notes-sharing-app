package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NotesDB []models.Notes
	result := db.Find(&NotesDB)
	if result.Error != nil {
		http.Error(w, "Database Error", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(NotesDB)
	w.WriteHeader(http.StatusOK)

}
