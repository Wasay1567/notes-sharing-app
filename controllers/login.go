package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("notes321")

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	if r.Body == nil {
		http.Error(w, "Request body missing", http.StatusBadRequest)
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u models.User
	result := db.Where("username = ? AND password = ?", payload.Username, payload.Password).First(&u)
	if result.Error != nil {
		http.Error(w, "Username NOT FOUND", http.StatusNotFound)
		return
	}

	token, err := GenerateJWT(payload.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
