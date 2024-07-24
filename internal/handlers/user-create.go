package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
}

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondWithBadRequest(w, "invalid request payload")
		return
	}

	id, err := c.userSrv.CreateUser(user.Username,
		user.Email, user.Password,
		user.Firstname, user.Lastname)
	if err != nil {
		msg := fmt.Errorf("failed to create user: %w", err).Error()
		respondWithInternalServerError(w, msg)
		return
	}

	response := CreateUserResponse{
		UUID: id.String(),
	}

	respondWithJSON(w, response, http.StatusCreated)
}
