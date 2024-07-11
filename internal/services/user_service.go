package services

import (
	"github.com/programme-lv/users-microservice/internal/entities"
	"github.com/programme-lv/users-microservice/internal/repository"
)

func CreateUser(user entities.User) error {
	return repository.CreateUser(user)
}

func GetUser(id string) (string, error) {
	user, err := repository.GetUser(id)
	if err != nil {
		return "", err
	}
	return user, nil
}

func UpdateUser(user entities.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(id string) error {
	return repository.DeleteUser(id)
}
