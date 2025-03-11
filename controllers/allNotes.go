package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func GetAllNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(models.NotesList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
