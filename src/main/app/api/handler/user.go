package handler

import (
	"encoding/json"
	"net/http"
)

// @Security BearerAuth
// @Summary Get users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r  *http.Request) {
	users, err := h.UserService.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(users)
	
}