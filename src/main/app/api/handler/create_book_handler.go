package handler

import (
	"app/src/main/app/model"
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Create book
// @Description Create new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body model.CreateBookModel true "Book data"
// @Success 201 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req model.CreateBookModel

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Decode error:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	book, err := h.CreateBookService.CreateBook(&req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}