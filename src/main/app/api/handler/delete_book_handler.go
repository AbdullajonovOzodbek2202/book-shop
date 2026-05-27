package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("ERROR invalid id:", err)

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid id",
		})

		return
	}

	err = h.DeleteBookService.DeleteBook(int64(id))
	if err != nil {
		log.Println("Error delete book:", err)

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to delete book",
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Book deleted successfully!",
	})
}