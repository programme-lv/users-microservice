package domain

import (
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID
	Username  string
	Email     string
	BcryptPwd []byte
}

func NewUser(uuid uuid.UUID, username string, email string, bcryptPwd []byte) (*User, error) {
	user := &User{
		UUID:      uuid,
		Username:  username,
		Email:     email,
		BcryptPwd: bcryptPwd,
	}
	return user, nil
}

func (u *User) Validate() (bool, string) {
	if u.Username == "" {
		return false, "username is required"
	}

	if u.Email == "" {
		return false, "email is required"
	}

	if len(u.Email) > 64 {
		return false, "email too long"
	}

	if len(u.Email) < 3 {
		return false, "email too short"
	}

	if len(u.Username) > 24 {
		return false, "username too long"
	}

	if len(u.Username) < 3 {
		return false, "username too short"
	}

	if !validEmail(u.Email) {
		return false, "invalid email"
	}

	return true, ""
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
