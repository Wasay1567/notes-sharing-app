package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"github.com/golang-jwt/jwt/v5"
)

var mu sync.Mutex

func CreateNewNote(w http.ResponseWriter, r *http.Request) {
	var payload models.GetNote

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	userId := GetUserID(r)
	if userId == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Println(userId)

	note := models.Notes{
		Title:      payload.Title,
		Content:    payload.Content,
		Expiration: payload.Expiration,
		Created_at: time.Now(),
		Viewed:     false,
		UserID:     userId,
		Password:   payload.Password,
	}

	err = validate(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&note)

	timeParts := strings.Split(note.Expiration, " ")
	if len(timeParts) == 2 && timeParts[1] != "view" {
		go deleteNote(timeParts[0], timeParts[1], &note)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Resp{
		Id:  note.Id,
		Url: "https://localhost:8080/v1/notes/" + strconv.Itoa(int(note.Id)),
	})
}

func GetUserID(r *http.Request) uint {
	tokenString := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	username := claims["username"].(string)
	var u models.User
	_ = db.Where("username = ?", username).First(&u)
	return u.ID
}

func validate(note *models.Notes) error {

	if note.Title == "" {
		return errors.New("title is missing")
	}
	if note.Expiration == "" {
		note.Expiration = "1 h"
	} else {
		n := strings.Split(note.Expiration, " ")
		if len(n) != 2 {
			return errors.New("invalid expiration format")
		}
	}
	return nil
}

func deleteNote(time_str, time_frame string, note *models.Notes) {
	t, _ := strconv.Atoi(time_str)
	switch time_frame {
	case "h":
		time.Sleep(time.Duration(t) * time.Hour)
	case "min":
		time.Sleep(time.Duration(t) * time.Minute)
	case "s":
		time.Sleep(time.Duration(t) * time.Second)
	}

	mu.Lock()
	defer mu.Unlock()
	// results := db.Where("id = ?", note.Id).Delete(&models.Notes{})
	// if results.Error != nil {
	// 	log.Fatal("Database Error", results.Error)
	// 	return
	// }
	// if results.RowsAffected == 0 {
	// 	fmt.Printf("Note with the id %v not found\n", note.Id)
	// 	return
	// }
	note.IsExpired = true
	db.Save(&note)
	fmt.Printf("Note expires with the id : %v\n", note.Id)
}
