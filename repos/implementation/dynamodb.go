package implementation

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getNewClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}

func insert(value interface{}, tableName string) error {
	dynamodbClient := getNewClient()
	attributeValue, err := dynamodbattribute.MarshalMap(value)

	if err != nil {
		fmt.Println(err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      attributeValue,
	}

	_, err2 := dynamodbClient.PutItem(input)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	return nil
}
