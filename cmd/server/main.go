package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/repository"
	"github.com/programme-lv/users-microservice/internal/service"
)

func main() {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		panic("JWT_KEY environment variable is not set")
	}

	userService := service.NewUserService(getDynamoDbRepo())
	controller := handlers.NewController(userService, []byte(jwtKey))

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller.RegisterRoutes(r)

	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", r)
}

func getDynamoDbRepo() *repository.DynamoDBUserRepository {
	tableName := os.Getenv("USERS_TABLE_NAME")
	if tableName == "" {
		panic("USERS_TABLE_NAME environment variable is not set")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"))
	if err != nil {
		panic(fmt.Sprintf("unable to load SDK config, %v", err))
	}
	dynamoClient := dynamodb.NewFromConfig(cfg)
	repo, err := repository.NewDynamoDBUserRepository(dynamoClient,
		tableName)
	if err != nil {
		panic(fmt.Sprintf("unable to create user repository, %v",
			err))
	}
	return repo
}
