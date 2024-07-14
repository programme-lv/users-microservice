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
	firstname string
	lastname  string
}

type UsernameUniquenessChecker interface {
	DoesUsernameExist(username string) (bool, error)
}

type EmailUniquenessChecker interface {
	DoesEmailExist(email string) (bool, error)
}

func NewUser(uuid uuid.UUID, username, email, password, firstname, lastname string,
	usernameUniquenessChecker UsernameUniquenessChecker,
	emailUniquenessChecker EmailUniquenessChecker) (User, error) {
	user := User{
		id: uuid,
	}

	var err error

	err = user.SetUsername(username, usernameUniquenessChecker)
	if err != nil {
		return User{}, err
	}

	err = user.SetEmail(email, emailUniquenessChecker)
	if err != nil {
		return User{}, err
	}

	err = user.SetPassword(password)
	if err != nil {
		return User{}, err
	}

	user.firstname = firstname
	user.lastname = lastname

	return user, nil
}

func RecoverUser(uuid uuid.UUID, username, email, bcryptPwd, firstname, lastname string) User {
	return User{
		id:        uuid,
		username:  username,
		email:     email,
		bcryptPwd: bcryptPwd,
		firstname: firstname,
		lastname:  lastname,
	}
}

func (u *User) SetUUID(uuid uuid.UUID) {
	u.id = uuid
}

func (u *User) SetUsername(username string,
	usernameUniquenessChecker UsernameUniquenessChecker) error {
	if username == "" {
		return errors.New("username is required")
	}

	if len(username) > 24 {
		return errors.New("username cannot be longer than 24 characters")
	}

	if len(username) < 3 {
		return errors.New("username cannot be shorter than 3 characters")
	}

	exists, err := usernameUniquenessChecker.DoesUsernameExist(username)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user with such username already exists")
	}

	u.username = username
	return nil
}

func (u *User) SetEmail(email string, emailUniquenessChecker EmailUniquenessChecker) error {
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

	exists, err := emailUniquenessChecker.DoesEmailExist(email)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("user with such email already exists")
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

func (u *User) GetFirstname() string {
	return u.firstname
}

func (u *User) GetLastname() string {
	return u.lastname
}
	_, err := mail.ParseAddress(email)
	return err == nil
}
