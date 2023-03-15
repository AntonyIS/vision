package database

import "github.com/aws/aws-sdk-go/service/dynamodb"

type dynamoDB struct {
	client                              *dynamodb.DynamoDB
	userTable, clientTable, deviceTable string
}
