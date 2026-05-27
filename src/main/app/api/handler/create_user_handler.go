package handler

import (
	"app/src/main/app/model"
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Create user
// @Description Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "User data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Router /createUser [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req model.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Decode error:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	user, err := h.CreateUserService.CreateUser(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}