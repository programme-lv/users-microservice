package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			lambdaHandler(controller.CreateUser, w, r)
		})
		r.Get("/{uuid}", func(w http.ResponseWriter, r *http.Request) {
			lambdaHandler(controller.GetUser, w, r)
		})
		r.Put("/{uuid}", func(w http.ResponseWriter, r *http.Request) {
			lambdaHandler(controller.UpdateUser, w, r)
		})
		r.Delete("/{uuid}", func(w http.ResponseWriter, r *http.Request) {
			lambdaHandler(controller.DeleteUser, w, r)
		})
	})

	http.ListenAndServe(":8080", r)
}
