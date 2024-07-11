package service

import (
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
	"github.com/programme-lv/users-microservice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.DynamoDBUserRepository
}

func NewUserService(repo repository.DynamoDBUserRepository) *UserService {
	return &UserService{repo: repo}
}

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

func (s *UserService) UpdateUser(input UpdateUserInput) error {
	// Implement the update logic here
	return nil
}

func (s *UserService) DeleteUser(uuid uuid.UUID) error {
	return s.repo.DeleteUser(uuid)
}
