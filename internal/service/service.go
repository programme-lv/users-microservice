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
