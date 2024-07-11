package repository

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/programme-lv/users-microservice/internal/domain"
)

type DynamoDBUserRepository struct {
	db        *dynamodb.Client
	tableName string
}

// NewDynamoDBUserRepository creates a new DynamoDBUserRepository with a DynamoDB client
func NewDynamoDBUserRepository(tableName string) (*DynamoDBUserRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		return nil, errors.New("unable to load SDK config, " + err.Error())
	}
	db := dynamodb.NewFromConfig(cfg)
	return &DynamoDBUserRepository{db: db, tableName: tableName}, nil
}

func (r *DynamoDBUserRepository) GetUser(uuid uuid.UUID) (domain.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: uuid.String()},
		},
	}

	result, err := r.db.GetItem(context.TODO(), input)
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

func (r *DynamoDBUserRepository) StoreUser(user domain.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      item,
	}

	_, err = r.db.PutItem(context.TODO(), input)
	return err
}

func (r *DynamoDBUserRepository) DeleteUser(uuid uuid.UUID) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: uuid.String()},
		},
	}

	_, err := r.db.DeleteItem(context.TODO(), input)
	return err
}

func (r *DynamoDBUserRepository) ListUsers() ([]domain.User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(r.tableName),
	}

	result, err := r.db.Scan(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = attributevalue.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
