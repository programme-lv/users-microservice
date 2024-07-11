package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Controller) CreateUser(ctx context.Context, request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	var user CreateUserRequest
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}
	}

	err = c.UserService.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusCreated}
}
