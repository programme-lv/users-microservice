package domain

import (
	"errors"
	"log/slog"
	"net/mail"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id        uuid.UUID
	username  string
	email     string
	bcryptPwd string
}

func NewUser(uuid uuid.UUID, username, email, password string) (User, error) {
	user := User{
		id: uuid,
	}

	var err error

	err = user.SetUsername(username)
	if err != nil {
		return User{}, err
	}

	err = user.SetEmail(email)
	if err != nil {
		return User{}, err
	}

	err = user.SetPassword(password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func RecoverUser(uuid uuid.UUID, username, email, bcryptPwd string) User {
	return User{
		id:        uuid,
		username:  username,
		email:     email,
		bcryptPwd: bcryptPwd,
	}
}

func (u *User) SetUUID(uuid uuid.UUID) {
	u.id = uuid
}

func (u *User) SetUsername(username string) error {
	if username == "" {
		return errors.New("username is required")
	}

	if len(username) > 24 {
		return errors.New("username cannot be longer than 24 characters")
	}

	if len(username) < 3 {
		return errors.New("username cannot be shorter than 3 characters")
	}

	u.username = username
	return nil
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	if len(email) > 64 {
		return errors.New("email cannot be longer than 64 characters")
	}

	if len(email) < 3 {
		return errors.New("email cannot be shorter than 3 characters")
	}

	if !validEmail(email) {
		return errors.New("invalid email")
	}

	u.email = email
	return nil
}

func (u *User) SetPassword(password string) error {
	bcryptPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		return err
	}

	u.bcryptPwd = string(bcryptPwd)
	return nil
}

func (u *User) GetUUID() uuid.UUID {
	return u.id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetBcryptPwd() string {
	return u.bcryptPwd
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
