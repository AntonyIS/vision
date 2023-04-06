package dynamodb

import (
	"errors"
	"fmt"

	"github.com/AntonyIS/vision/config"
	"github.com/AntonyIS/vision/internal/core/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type dynamoDbClient struct {
	client                              *dynamodb.DynamoDB
	userTable, clientTable, deviceTable string
}

func NewDynamoDB(c *config.BaseConfig) *dynamoDbClient {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(c.AWS_DEFAULT_REGION),
	}))
	return &dynamoDbClient{
		client:      dynamodb.New(sess),
		userTable:   c.DynamoDB_UsersTable,
		deviceTable: c.DynamoDB_DevicesTable,
		clientTable: c.DynamoDB_ClientsTable,
	}

}

func (db *dynamoDbClient) CreateUser(user *domain.User) (*domain.User, error) {

	entityParsed, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.userTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	user, err = db.ReadUser(user.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *dynamoDbClient) ReadUser(id string) (*domain.User, error) {
	result, err := db.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.userTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New(fmt.Sprintf("user with id [ %s ] not found", id))
	}
	var user domain.User
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *dynamoDbClient) ReadUsers() ([]*domain.User, error) {
	users := []*domain.User{}
	filt := expression.Name("Id").AttributeNotExists()
	proj := expression.NamesList(
		expression.Name("user_id"),
		expression.Name("firstname"),
		expression.Name("lastname"),
		expression.Name("email"),
		expression.Name("phone"),
		expression.Name("password"),
		expression.Name("image"),
		expression.Name("client"),
		expression.Name("devices"),
		expression.Name("is_admin"),
		expression.Name("is_super_user"),
		expression.Name("is_installer"),
		expression.Name("is_client"),
		expression.Name("is_manager"),
		expression.Name("is_technician"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(db.userTable),
	}
	result, err := db.client.Scan(params)

	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		var user domain.User

		err = dynamodbattribute.UnmarshalMap(item, &user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)

	}

	return users, nil
}

func (db *dynamoDbClient) UpdateUser(user *domain.User) (*domain.User, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.userTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *dynamoDbClient) DeleteUser(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(db.userTable),
	}

	res, err := db.client.DeleteItem(input)
	if res == nil {
		return errors.New(fmt.Sprintf("No user to delete: %v", err))
	}
	if err != nil {
		return errors.New(fmt.Sprintf("Got error calling DeleteItem: %v", err))
	}
	return nil
}

func (db *dynamoDbClient) CreateDevice(device *domain.Device) (*domain.Device, error) {

	entityParsed, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.deviceTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (db *dynamoDbClient) ReadDevice(id string) (*domain.Device, error) {
	result, err := db.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.deviceTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New(fmt.Sprintf("device with id [ %s ] not found", id))
	}
	var device domain.Device
	err = dynamodbattribute.UnmarshalMap(result.Item, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (db *dynamoDbClient) ReadDevices() ([]*domain.Device, error) {
	devices := []*domain.Device{}
	filt := expression.Name("Id").AttributeNotExists()
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("title"),
		expression.Name("body"),
		expression.Name("user_id"),
		expression.Name("created_at"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(db.deviceTable),
	}
	result, err := db.client.Scan(params)
	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		var device domain.Device

		err = dynamodbattribute.UnmarshalMap(item, &device)
		if err != nil {
			return nil, err
		}
		devices = append(devices, &device)

	}

	return devices, nil
}

func (db *dynamoDbClient) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.deviceTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return device, nil
}

func (db *dynamoDbClient) DeleteDevice(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(db.deviceTable),
	}

	res, err := db.client.DeleteItem(input)
	if res == nil {
		return errors.New(fmt.Sprintf("no project to delete: %s", err))
	}
	if err != nil {
		return errors.New(fmt.Sprintf("got error calling DeleteItem: %s", err))
	}
	return nil
}

func (db *dynamoDbClient) CreateClient(client *domain.Client) (*domain.Client, error) {

	entityParsed, err := dynamodbattribute.MarshalMap(client)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.clientTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (db *dynamoDbClient) ReadClient(id string) (*domain.Client, error) {
	result, err := db.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.clientTable),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New(fmt.Sprintf("client with id [ %s ] not found", id))
	}
	var client domain.Client
	err = dynamodbattribute.UnmarshalMap(result.Item, &client)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (db *dynamoDbClient) ReadClients() ([]*domain.Client, error) {
	clients := []*domain.Client{}
	filt := expression.Name("Id").AttributeNotExists()
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("title"),
		expression.Name("body"),
		expression.Name("user_id"),
		expression.Name("created_at"),
	)
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(db.clientTable),
	}
	result, err := db.client.Scan(params)
	if err != nil {
		return nil, err
	}

	for _, item := range result.Items {
		var client domain.Client

		err = dynamodbattribute.UnmarshalMap(item, &client)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client)

	}

	return clients, nil
}

func (db *dynamoDbClient) UpdateClient(client *domain.Client) (*domain.Client, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(client)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.clientTable),
	}

	_, err = db.client.PutItem(input)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (db *dynamoDbClient) DeleteClient(id string) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(db.clientTable),
	}

	res, err := db.client.DeleteItem(input)
	if res == nil {
		return errors.New(fmt.Sprintf("no client to delete: %s", err))
	}
	if err != nil {
		return errors.New(fmt.Sprintf("got error calling DeleteItem: %s", err))
	}
	return nil
}
