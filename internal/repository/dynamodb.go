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

func mapDomainUserToDynamoDBUser(user domain.User) map[string]interface{} {
	return map[string]interface{}{
		"uuid":       user.UUID.String(),
		"username":   user.Username,
		"email":      user.Email,
		"bcrypt_pwd": string(user.BcryptPwd),
	}
}

func mapDynamoDBUserToDomainUser(dict map[string]interface{}) (domain.User, error) {
	_, uuidFound := dict["uuid"]
	_, usernameFound := dict["username"]
	_, emailFound := dict["email"]
	_, bcryptPwdFound := dict["bcrypt_pwd"]

	if !uuidFound || !usernameFound || !emailFound || !bcryptPwdFound {
		return domain.User{}, errors.New("missing fields")
	}

	uuid, err := uuid.Parse(dict["uuid"].(string))
	if err != nil {
		return domain.User{}, errors.New("error parsing UUID")
	}

	username, ok := dict["username"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid username")
	}

	email, ok := dict["email"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid email")
	}

	bcryptPwd, ok := dict["bcrypt_pwd"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid bcrypt password")
	}

	return domain.User{
		UUID:      uuid,
		Username:  username,
		Email:     email,
		BcryptPwd: []byte(bcryptPwd),
	}, nil
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

func (r *DynamoDBUserRepository) GetUser(id uuid.UUID) (domain.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"uuid": &types.AttributeValueMemberS{Value: id.String()},
		},
	}

	result, err := r.db.GetItem(context.TODO(), input)
	if err != nil {
		return domain.User{}, err
	}

	if result.Item == nil {
		return domain.User{}, errors.New("user not found")
	}

	dict := map[string]interface{}{}

	err = attributevalue.UnmarshalMap(result.Item, &dict)
	if err != nil {
		return domain.User{}, err
	}

	user, err := mapDynamoDBUserToDomainUser(dict)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *DynamoDBUserRepository) StoreUser(user domain.User) error {
	record := mapDomainUserToDynamoDBUser(user)

	item, err := attributevalue.MarshalMap(record)
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
