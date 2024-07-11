package main

import (
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/repository"
	"github.com/programme-lv/users-microservice/internal/service"
)

func main() {
	repo, err := repository.NewDynamoDBUserRepository()
	if err != nil {
		panic(err)
	}
	userService := service.NewUserService(repo)
	controller := handlers.NewController(userService)

	// Example of how to start the lambda with a specific handler
	// lambda.Start(controller.CreateUser)
}
