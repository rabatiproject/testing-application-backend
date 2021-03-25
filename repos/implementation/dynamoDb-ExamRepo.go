package implementation

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

type dynamoDbExamRepo struct {
	tableName string
}

func NewDynamoDbRepo() *dynamoDbExamRepo {
	return &dynamoDbExamRepo{
		tableName: "EXAMS",
	}
}

func getNewClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}

func (d *dynamoDbExamRepo) SaveExam(exam *base.Exam) error {
	dynamodbClient := getNewClient()
	exam.Id = uuid.New().String()
	attributeValue, err := dynamodbattribute.MarshalMap(exam)

	if err != nil {
		fmt.Println("hataaa")
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(NewDynamoDbRepo().tableName),
	}

	_, err2 := dynamodbClient.PutItem(input)
	if err2 != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err2.Error())
	}
	return nil
}
