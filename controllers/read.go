package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract and validate the ID from the URL path
	id := r.PathValue("id")
	pass := r.PathValue("pass")

	var note models.Notes
	result := db.First(&note, "id = ? AND password = ?", id, pass)
	if result.Error != nil {
		http.Error(w, "ID NOT FOUND", http.StatusNotFound)
	}

	err := json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusFound)

}
