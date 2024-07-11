package domain

import (
	"errors"
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

func (u *User) Validate() error {
	if len(u.Email) > 64 {
		return errors.New("email too long")
	}

	if len(u.Email) < 3 {
		return errors.New("email too short")
	}

	if len(u.Username) > 24 {
		return errors.New("username too long")
	}

	if len(u.Username) < 3 {
		return errors.New("username too short")
	}

	if !validEmail(u.Email) {
		return errors.New("invalid email")
	}

	return nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
