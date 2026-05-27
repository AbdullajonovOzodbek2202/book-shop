package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Get Books
// @Description Get all books
// @Tags books
// @Produce json
// @Success 200 {array} model.Book
// @Router /books [get]
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.GetBookService.GetBooks()
	if err != nil {
		log.Println("ERROR:", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		log.Println("ENCODE ERROR:", err)
	}
}