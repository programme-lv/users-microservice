package handlers

import (
	"github.com/programme-lv/users-microservice/internal/service"
)

type Controller struct {
	UserService *service.UserService
}

func NewController(userService *service.UserService) *Controller {
	return &Controller{UserService: userService}
}
