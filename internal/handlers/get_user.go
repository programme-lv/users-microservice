package handlers

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

type GetUserResponse struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (c *Controller) GetUser(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uuidParam := request.PathParameters["uuid"]
	id, err := uuid.Parse(uuidParam)
	if err != nil {
		return respondWithBadRequest("Invalid UUID"), nil
	}

	user, err := c.UserService.GetUser(id)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusNotFound}, err
	}

	response := GetUserResponse{
		UUID:     user.UUID.String(),
		Username: user.Username,
		Email:    user.Email,
	}

	return respondWithJSON(response)
}
