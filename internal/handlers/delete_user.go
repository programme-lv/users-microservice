package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")

	id, err := uuid.Parse(uuidParam)
	if err != nil {
		respondWithBadRequest(w, "Invalid UUID")
		return
	}

	err = c.UserService.DeleteUser(id)
	if err != nil {
		respondWithInternalServerError(w, "Failed to delete user")
		return
	}

	w.WriteHeader(http.StatusOK)
}
