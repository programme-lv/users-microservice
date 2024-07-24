package handlers

import (
	"net/http"
)

type ListUsersResponse struct {
	Users []UserResponse `json:"users"`
}

type UserResponse struct {
	UUID      string  `json:"uuid"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Firstname *string `json:"firstname,omitempty"`
	Lastname  *string `json:"lastname,omitempty"`
}

func (c *Controller) ListUsers(w http.ResponseWriter, r *http.Request) {
	user, err := c.UserService.ListUsers()
	if err != nil {
		respondWithBadRequest(w, "user not found")
		return
	}

	users := []UserResponse{}
	for _, u := range user {
		users = append(users, UserResponse{
			UUID:      u.GetUUID().String(),
			Username:  u.GetUsername(),
			Email:     u.GetEmail(),
			Firstname: u.GetFirstname(),
			Lastname:  u.GetLastname(),
		})
	}

	response := ListUsersResponse{
		Users: users,
	}

	respondWithJSON(w, response, http.StatusOK)
}
