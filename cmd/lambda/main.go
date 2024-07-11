package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/repository"
	"github.com/programme-lv/users-microservice/internal/service"
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

	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUser)
		r.Get("/{uuid}", controller.GetUser)
		r.Put("/{uuid}", controller.UpdateUser)
		r.Delete("/{uuid}", controller.DeleteUser)
	})

	http.ListenAndServe(":8080", r)
}
