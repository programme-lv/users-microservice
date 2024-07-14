package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/handlers"
	"github.com/programme-lv/users-microservice/internal/service"
)

func main() {
	userService := service.NewUserService(nil)
	controller := handlers.NewController(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

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
			r.Post("/register", controller.CreateUser)
			r.Post("/login", controller.LoginUser)
		})
	})

	http.ListenAndServe(":8080", r)
}
