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

	result := db.Where("id = ? AND password = ?", id, pass).Delete(&models.Notes{})
	if result.Error != nil {
		http.Error(w, "DB ERROR", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Note not found or password is incorrect", http.StatusNotFound)
	}
	w.Write([]byte("Note Deleted"))
	w.WriteHeader(http.StatusOK)
}
