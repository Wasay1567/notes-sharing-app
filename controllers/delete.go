package controllers

import (
	"net/http"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pass := r.PathValue("pass")
	mu.Lock()
	defer mu.Unlock()
	for i, note := range models.NotesList {
		if note.Id == id {
			if note.Password == pass {
				models.NotesList = append(models.NotesList[:i], models.NotesList[i+1:]...)
				w.Write([]byte("Note deleted"))
				w.WriteHeader(http.StatusAccepted)
				return
			} else {
				http.Error(w, "Incorrect Password", http.StatusForbidden)
				return
			}
		}
	}

	http.Error(w, "ID NOT FOUND", http.StatusNotFound)
}
