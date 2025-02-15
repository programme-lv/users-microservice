package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type GetUserResponse struct {
	UUID      string  `json:"uuid"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(uuidParam)
	if err != nil {
		respondWithBadRequest(w, "invalid UUID")
		return
	}

	user, err := c.userSrv.GetUser(id)
	if err != nil {
		respondWithBadRequest(w, "user not found")
		return
	}

	response := GetUserResponse{
		UUID:      user.GetUUID().String(),
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		Firstname: user.GetFirstname(),
		Lastname:  user.GetLastname(),
	}

	respondWithJSON(w, response, http.StatusOK)
}
