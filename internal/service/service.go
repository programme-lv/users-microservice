package service

import (
	"github.com/programme-lv/users-microservice/internal/repository"
)

type UserService struct {
	repo *repository.DynamoDBUserRepository
}

func NewUserService(repo *repository.DynamoDBUserRepository) *UserService {
	return &UserService{repo: repo}
}
	"golang.org/x/crypto/bcrypt"
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
