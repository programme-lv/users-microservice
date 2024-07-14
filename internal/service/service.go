package service

import (
	"errors"

	"github.com/programme-lv/users-microservice/internal/domain"
	"github.com/programme-lv/users-microservice/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.DynamoDBUserRepository
}

func NewUserService(repo *repository.DynamoDBUserRepository) *UserService {
	return &UserService{repo: repo}
}
func (s *UserService) AuthenticateUser(username, password string) (domain.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.GetBcryptPwd()), []byte(password))
	if err != nil {
		return domain.User{}, errors.New("invalid password")
	}

	return user, nil
}
