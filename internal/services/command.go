package services

import (
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/repository"
)

type UpdateUserInput struct {
	UUID uuid.UUID
}

func CreateUser() error {
	return nil
}

func UpdateUser(input UpdateUserInput) error {
	return nil
}

func DeleteUser(uuid uuid.UUID) error {
	return repository.DeleteUser(uuid)
}
