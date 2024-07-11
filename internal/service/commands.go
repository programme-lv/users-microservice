package service

import (
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) CreateUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		UUID:      uuid.New(),
		Username:  username,
		Email:     email,
		BcryptPwd: hashedPassword,
	}

	return s.repo.StoreUser(user)
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

	return s.repo.StoreUser(user)
}

func (s *UserService) DeleteUser(uuid uuid.UUID) error {
	return s.repo.DeleteUser(uuid)
}
