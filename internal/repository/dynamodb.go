package repository

import (
	"context"
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
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

func (r *DynamoDBUserRepository) NewEmailUniquenessChecker() domain.EmailUniquenessChecker {
	return &dynamoDBEmailUniquenessChecker{repo: r}
}

type dynamoDBEmailUniquenessChecker struct {
	repo *DynamoDBUserRepository
}

// DoesEmailExist implements domain.EmailUniquenessChecker.
func (d *dynamoDBEmailUniquenessChecker) DoesEmailExist(email string) (bool, error) {
	// for now just list all users, iterate through
	users, err := d.repo.ListUsers()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.GetEmail() == email {
			return true, nil
		}
	}

	return false, nil
}

func IsNilish(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)
	k := v.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer,
		reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return false
}

func mapDomainUserToDynamoDBUser(user domain.User) map[string]interface{} {
	return map[string]interface{}{
		"uuid":       user.GetUUID().String(),
		"username":   user.GetUsername(),
		"email":      user.GetEmail(),
		"bcrypt_pwd": user.GetBcryptPwd(),
		"firstname":  user.GetFirstname(),
		"lastname":   user.GetLastname(),
	}
}

func mapDynamoDBUserToDomainUser(dict map[string]interface{}) (domain.User, error) {
	_, uuidFound := dict["uuid"]
	_, usernameFound := dict["username"]
	_, emailFound := dict["email"]
	_, bcryptPwdFound := dict["bcrypt_pwd"]
	_, firstnameFound := dict["firstname"]
	_, lastnameFound := dict["lastname"]

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

	var firstname *string = nil
	if firstnameFound && !IsNilish(dict["firstname"]) {
		firstnameStr, ok := dict["firstname"].(string)
		if !ok {
			return domain.User{}, errors.New("invalid firstname")
		}
		firstname = &firstnameStr
	}

	var lastname *string = nil
	if lastnameFound && !IsNilish(dict["lastname"]) {
		lastnameStr, ok := dict["lastname"].(string)
		if !ok {
			return domain.User{}, errors.New("invalid lastname")
		}
		lastname = &lastnameStr
	}

	return domain.ConstructUser(uuid, username, email, bcryptPwd, firstname, lastname), nil
}

// NewDynamoDBUserRepository creates a new DynamoDBUserRepository with a DynamoDB client
func NewDynamoDBUserRepository(db *dynamodb.Client, tableName string) (*DynamoDBUserRepository, error) {
	return &DynamoDBUserRepository{db: db, tableName: tableName}, nil
}

func (r *DynamoDBUserRepository) GetUserByUsername(username string) (domain.User, error) {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(r.tableName),
		FilterExpression: aws.String("username = :username"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":username": &types.AttributeValueMemberS{Value: username},
		},
	}

	result, err := r.db.Scan(context.TODO(), input)
	if err != nil {
		return domain.User{}, err
	}

	if len(result.Items) == 0 {
		return domain.User{}, errors.New("user not found")
	}

	dict := map[string]interface{}{}
	err = attributevalue.UnmarshalMap(result.Items[0], &dict)
	if err != nil {
		return domain.User{}, err
	}

	return mapDynamoDBUserToDomainUser(dict)
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

	var dicts []map[string]interface{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &dicts)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, dict := range dicts {
		user, err := mapDynamoDBUserToDomainUser(dict)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *DynamoDBUserRepository) NewUsernameUniquenessChecker() domain.UsernameUniquenessChecker {
	return &dynamoDBUsernameUniquenessChecker{repo: r}
}

type dynamoDBUsernameUniquenessChecker struct {
	repo *DynamoDBUserRepository
}

// DoesUsernameExist implements domain.UsernameUniquenessChecker.
func (d *dynamoDBUsernameUniquenessChecker) DoesUsernameExist(username string) (bool, error) {
	// for now just list all users, iterate through
	users, err := d.repo.ListUsers()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.GetUsername() == username {
			return true, nil
		}
	}

	return false, nil
}
