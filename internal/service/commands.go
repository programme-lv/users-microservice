package service

import (
	"errors"
	"log/slog"

	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(username, email, password string) error {
	slog.Info("Creating user", "username", username, "email", email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("Failed to hash password", "error", err)
		return err
	}

	user := domain.User{
		UUID:      uuid.New(),
		Username:  username,
		Email:     email,
		BcryptPwd: hashedPassword,
	}

	ok, msg := user.Validate()
	if !ok {
		return errors.New(msg)
	}

	err = s.repo.StoreUser(user)
	if err != nil {
		slog.Error("Failed to store user", "error", err)
		return err
	}
	return err
}

type UpdateUserInput struct {
	UUID     uuid.UUID
	Username *string
	Email    *string
}

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	user, err := s.repo.GetUser(input.UUID)
	if err != nil {
		return err
	}

	if input.Username != nil {
		user.Username = *input.Username
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	ok, msg := user.Validate()
	if !ok {
		return errors.New(msg)
	}

	return s.repo.StoreUser(user)
}

func (s *UserService) DeleteUser(uuid uuid.UUID) error {
	return s.repo.DeleteUser(uuid)
}
