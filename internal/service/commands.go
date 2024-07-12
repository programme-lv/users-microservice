package service

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
)

func (s *UserService) CreateUser(username, email, password string) (uuid.UUID, error) {
	slog.Info("Creating user", "username", username, "email", email)

	user, err := domain.NewUser(uuid.New(), username, email, password,
		s.repo.NewUsernameUniquenessChecker())
	if err != nil {
		return uuid.Nil, err
	}

	err = s.repo.StoreUser(user)
	if err != nil {
		slog.Error("Failed to store user", "error", err)
		return uuid.Nil, err
	}
	return user.GetUUID(), nil
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
		err = user.SetUsername(*input.Username,
			s.repo.NewUsernameUniquenessChecker())
		if err != nil {
			return err
		}
	}

	if input.Email != nil {
		err = user.SetEmail(*input.Email)
		if err != nil {
			return err
		}
	}

	return s.repo.StoreUser(user)
}

func (s *UserService) DeleteUser(uuid uuid.UUID) error {
	return s.repo.DeleteUser(uuid)
}
