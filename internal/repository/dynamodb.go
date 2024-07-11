package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/programme-lv/users-microservice/internal/entities"
)

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-west-2"))

func CreateUser(user entities.User) error {
	// Implement DynamoDB logic to create a user
	return nil
}

func GetUser(id string) (entities.User, error) {
	// Implement DynamoDB logic to get a user by ID
	return entities.User{}, nil
}

func UpdateUser(user entities.User) error {
	// Implement DynamoDB logic to update a user
	return nil
}

func DeleteUser(id string) error {
	// Implement DynamoDB logic to delete a user by ID
	return nil
}
