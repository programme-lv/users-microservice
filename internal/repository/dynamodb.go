package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/programme-lv/users-microservice/internal/entities"
)

var db *dynamodb.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	db = dynamodb.NewFromConfig(cfg)
}

func CreateUser(user entities.User) error {
	// Implement DynamoDB logic to create a user
	_, err := db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Users"),
		Item: map[string]types.AttributeValue{
			"ID":    &types.AttributeValueMemberS{Value: user.ID},
			"Name":  &types.AttributeValueMemberS{Value: user.Name},
			"Email": &types.AttributeValueMemberS{Value: user.Email},
		},
	})
	return err
}

func GetUser(id string) (entities.User, error) {
	// Implement DynamoDB logic to get a user by ID
	result, err := db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		ID:    result.Item["ID"].(*types.AttributeValueMemberS).Value,
		Name:  result.Item["Name"].(*types.AttributeValueMemberS).Value,
		Email: result.Item["Email"].(*types.AttributeValueMemberS).Value,
	}
	return user, nil
}

func UpdateUser(user entities.User) error {
	// Implement DynamoDB logic to update a user
	_, err := db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: user.ID},
		},
		ExpressionAttributeNames: map[string]string{
			"#N": "Name",
			"#E": "Email",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":n": &types.AttributeValueMemberS{Value: user.Name},
			":e": &types.AttributeValueMemberS{Value: user.Email},
		},
		UpdateExpression: aws.String("SET #N = :n, #E = :e"),
	})
	return err
}

func DeleteUser(id string) error {
	// Implement DynamoDB logic to delete a user by ID
	_, err := db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Users"),
		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{Value: id},
		},
	})
	return err
}
