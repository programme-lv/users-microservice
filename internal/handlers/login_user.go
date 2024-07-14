package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/programme-lv/users-microservice/internal/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithBadRequest(w, "invalid request payload")
		return
	}

	user, err := c.UserService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		respondWithBadRequest(w, "invalid username or password")
		return
	}

	token, err := auth.GenerateJWT(user.GetUsername())
	if err != nil {
		respondWithInternalServerError(w, "failed to generate token")
		return
	}

	response := LoginResponse{Token: token}
	respondWithJSON(w, response, http.StatusOK)
}
