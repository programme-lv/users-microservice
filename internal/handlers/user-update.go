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
		respondWithBadRequest(w, "Invalid request payload")
		return
	}

	err = c.userSrv.UpdateUser(service.UpdateUserInput{
		UUID: [16]byte{},
	})
	if err != nil {
		respondWithInternalServerError(w, "Failed to update user")
		return
	}

	w.WriteHeader(http.StatusOK)
}
