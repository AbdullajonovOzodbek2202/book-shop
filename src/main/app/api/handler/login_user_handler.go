package handler

import (
	"app/src/main/app/model"
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.LoginUserModel true "Login data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req model.LoginUserModel

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("ERROR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	token, err := h.LoginUserService.LoginUser(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
