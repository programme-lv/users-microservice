package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: uuid.String()},
		},
	}

	result, err := db.GetItem(context.TODO(), input)
	if err != nil {
		return domain.User{}, err
	}

	if result.Item == nil {
		return domain.User{}, errors.New("user not found")
	}

	var user domain.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func StoreUser(user domain.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item:      item,
	}

	_, err = db.PutItem(context.TODO(), input)
	return err
}

func DeleteUser(uuid uuid.UUID) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: uuid.String()},
		},
	}

	_, err := db.DeleteItem(context.TODO(), input)
	return err
}
