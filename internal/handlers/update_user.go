package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/programme-lv/users-microservice/internal/service"
)

type UpdateUserRequest struct {
	UUID     *string `json:"uuid"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

func (c *Controller) UpdateUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user UpdateUserRequest
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, err
	}

	err = c.UserService.UpdateUser(service.UpdateUserInput{
		UUID: [16]byte{},
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}, err
	}

	return events.APIGatewayProxyResponse{StatusCode: http.StatusOK}, nil
}
