package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
)

var db *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	db = dynamodb.NewFromConfig(cfg)
}

func GetUser(uuid uuid.UUID) (domain.User, error) {
	panic("not implemented")
}

func StoreUser(user domain.User) error {
	panic("not implemented")
}

func DeleteUser(uuid uuid.UUID) error {
	panic("not implemented")
}
