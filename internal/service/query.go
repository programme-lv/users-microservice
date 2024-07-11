package service

import (
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
	"github.com/programme-lv/users-microservice/internal/repository"
)

func GetUser(uuid uuid.UUID) (domain.User, error) {
	return repository.GetUser(uuid)
}
func ListUsers() ([]domain.User, error) {
	return repository.ListUsers()
}
