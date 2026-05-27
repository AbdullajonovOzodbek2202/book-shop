package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// @Summary Get book by ID
// @Description Get single user
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} model.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [get]
func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {

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

	book, err := h.GetBookByIdService.GetBook(int64(id))
	if err != nil {
		log.Println("ERROR get user:", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(book)

}
