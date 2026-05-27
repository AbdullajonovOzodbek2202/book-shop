package handler

import (
	"app/src/main/app/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// @Summary Update book
// @Description Update book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param user body model.UpdateBookModel true "Update book data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [put]
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {


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

	var req model.UpdateBookModel

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("ERROR decode:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = h.UpdateBookService.UpdateBook(int64(id), &req)
	if err != nil {
		log.Println("ERROR update:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Book updated successfully",
	})
}