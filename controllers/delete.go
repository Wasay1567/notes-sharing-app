package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AbdulWasay1207/notes-sharing-app/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}
	user_id := GetUserID(r)

	mu.Lock()
	defer mu.Unlock()

	result := db.Where("id = ? AND user_id = ?", id, user_id).Delete(&models.Notes{})
	if result.Error != nil {
		log.Printf("Database error: %v", result.Error)
		http.Error(w, "Cannot delete the note", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Note not found or unauthorized", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Note Deleted"))
}
