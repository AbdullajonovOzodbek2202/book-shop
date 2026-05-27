package handler

import (
	"app/src/main/app/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// @Summary Update user
// @Description Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.UpdateUserModel true "Update user data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {


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

	var req model.UpdateUserModel

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("ERROR decode:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = h.UpdateUserService.UpdateUser(int64(id), &req)
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
		"message": "user updated successfully",
	})
}