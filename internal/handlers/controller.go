package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/programme-lv/users-microservice/internal/service"
)

type Controller struct {
	userSrv *service.UserService
	jwtKey  []byte
}

func NewController(userService *service.UserService) *Controller {
	return &Controller{userSrv: userService}
}

func (c *Controller) RegisterRoutes(r chi.Router) {
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", c.ListUsers)
			r.Post("/", c.CreateUser)
			r.Get("/{uuid}", c.GetUser)
			r.Put("/{uuid}", c.UpdateUser)
			r.Delete("/{uuid}", c.DeleteUser)
		})
	})

	r.Route("/auth", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Get("/", c.AuthState)
			r.Post("/login", c.LoginUser)
			r.Post("/register", c.Register)
		})
	})

}
