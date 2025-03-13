package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid Format", http.StatusBadRequest)
		return
	}

	u := models.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	result := db.Create(&u)
	if result.Error != nil {
		http.Error(w, "Username Already Exists", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
