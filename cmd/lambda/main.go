package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/handlers/responses"
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
