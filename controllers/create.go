package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
	"github.com/sethvargo/go-password/password"
)

var mu sync.Mutex

func CreateNewNote(w http.ResponseWriter, r *http.Request) {
	var payload models.GetNote

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p, _ := password.Generate(6, 3, 0, false, false)

	mu.Lock()
	note := models.Notes{
		Id:         p,
		Title:      payload.Title,
		Content:    payload.Content,
		Expiration: payload.Expiration,
		Created_at: time.Now().Format("Mon Jan 2 15:04"),
		Password:   payload.Password,
	}
	note_c := models.Resp{
		Id:  note.Id,
		Url: "https://localhost:8080/v1/notes/" + note.Id,
	}
	mu.Unlock()

	err = validate(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	models.NotesList = append(models.NotesList, note)
	result := db.Create(&note)
	mu.Unlock()
	if result.Error != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	time := strings.Split(note.Expiration, " ")
	if time[1] != "view" {
		go deleteNote(time[0], time[1], &note)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note_c)
}

func validate(note *models.Notes) error {
	var results []models.Notes
	db.Where("id = ?", note.Id).Find(&results)

	for len(results) > 0 {
		p, _ := password.Generate(6, 3, 0, false, false)
		note.Id = p
		db.Where("id = ?", note.Id).Find(&results)
	}

	if note.Title == "" {
		return errors.New("title is missing")
	}
	if note.Expiration == "" {
		note.Expiration = "1 min"
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
	results := db.Where("id = ?", note.Id).Delete(&models.Notes{})
	if results.Error != nil {
		log.Fatal("Database Error", results.Error)
		return
	}
	if results.RowsAffected == 0 {
		fmt.Printf("Note with the id %v not found\n", note.Id)
		return
	}
	fmt.Printf("Note expires with the id : %v\n", note.Id)
}
