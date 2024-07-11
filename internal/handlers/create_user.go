package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithBadRequest(w, "invalid request payload")
		return
	}

	err = c.UserService.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		msg := fmt.Errorf("failed to create user: %w", err).Error()
		respondWithInternalServerError(w, msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
