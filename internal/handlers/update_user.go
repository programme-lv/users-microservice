package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/programme-lv/users-microservice/internal/service"
)

type UpdateUserRequest struct {
	UUID     *string `json:"uuid"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = c.UserService.UpdateUser(service.UpdateUserInput{
		UUID: [16]byte{},
	})
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
