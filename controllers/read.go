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

	// Search for the note with the given ID
	for _, note := range models.NotesList {
		if note.Id == id {
			if note.Password == pass {
				// Encode the note as JSON and send it in the response
				err := json.NewEncoder(w).Encode(note)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			} else {
				http.Error(w, "password is incorrect", http.StatusForbidden)
				return
			}
		}
	}

	// If the note is not found, return a 404 Not Found error
	http.Error(w, "ID NOT FOUND", http.StatusNotFound)
}
