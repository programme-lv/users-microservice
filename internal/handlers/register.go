package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/programme-lv/users-microservice/internal/auth"
)

type RegisterRequest struct {
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
}

type RegisterResponse struct {
	UUID  string `json:"uuid"`
	Token string `json:"token"`
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		respondWithBadRequest(w, "invalid request payload")
		return
	}

	id, err := c.UserService.CreateUser(request.Username,
		request.Email, request.Password,
		request.Firstname, request.Lastname)
	if err != nil {
		msg := fmt.Errorf("failed to create user: %w", err).Error()
		respondWithInternalServerError(w, msg)
		return
	}

	user, err := c.UserService.GetUser(id)
	if err != nil {
		msg := fmt.Errorf("failed to get user: %w", err).Error()
		respondWithInternalServerError(w, msg)
		return
	}

	token, err := auth.GenerateJWT(user.GetUsername(),
		user.GetEmail(), user.GetUUID().String(),
		user.GetFirstname(), user.GetLastname())
	if err != nil {
		respondWithInternalServerError(w, "failed to generate token")
		return
	}

	response := RegisterResponse{
		UUID:  id.String(),
		Token: token,
	}

	respondWithJSON(w, response, http.StatusCreated)
}
