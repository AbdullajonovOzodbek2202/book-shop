package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// @Summary Get user by ID
// @Description Get single user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := h.GetUserService.GetUser(int64(id))
	if err != nil {
		log.Println("ERROR get user:", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}