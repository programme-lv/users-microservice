package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/repository"
	"github.com/programme-lv/users-microservice/internal/service"

	awschi "github.com/awslabs/aws-lambda-go-api-proxy/chi"
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
	r.Use(corsMiddleware)

	controller.RegisterRoutes(r)

	chiLambda := awschi.NewV2(r)

	handler := func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (
		events.APIGatewayV2HTTPResponse, error) {
		return chiLambda.ProxyWithContextV2(ctx, req)
	}

	lambda.Start(handler)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
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
