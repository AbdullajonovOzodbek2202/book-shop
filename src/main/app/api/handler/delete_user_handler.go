package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

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

	err = h.DeleteUserService.DeleteUser(int64(id))
	if err != nil {

		log.Println("ERROR delete user:", err)

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted successfully",
	})
}