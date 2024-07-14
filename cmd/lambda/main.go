package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/repository"
	"github.com/programme-lv/users-microservice/internal/service"

	awschi "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

func main() {
	tableName := os.Getenv("TABLE_NAME")
	if tableName == "" {
		panic("TABLE_NAME environment variable is not set")
	}

	repo, err := repository.NewDynamoDBUserRepository(tableName)
	if err != nil {
		panic(err)
	}
	userService := service.NewUserService(repo)
	controller := handlers.NewController(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")

			// Handle preflight request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", controller.ListUsers)
			r.Post("/", controller.CreateUser)
			r.Get("/{uuid}", controller.GetUser)
			r.Put("/{uuid}", controller.UpdateUser)
			r.Delete("/{uuid}", controller.DeleteUser)
		})
	})

	r.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// register is handled by create user
			// r.Post("/register", controller.RegisterUser)
			r.Post("/login", controller.LoginUser)
			r.Post("/register", controller.Register)
		})
	})

	chiLambda := awschi.NewV2(r)

	handler := func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return chiLambda.ProxyWithContextV2(ctx, req)
	}

	lambda.Start(handler)
}
