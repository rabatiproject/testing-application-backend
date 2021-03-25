package implementation

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
)

type dynamoDbExamRepo struct {
	TableName string
}

func NewDynamoDbRepo() *dynamoDbExamRepo {
	return &dynamoDbExamRepo{
		TableName: "EXAMS",
	}
}

func getNewClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return dynamodb.New(sess)
}

func (d *dynamoDbExamRepo) SaveExam(exam *base.Exam) error {
	exam.Id = uuid.New().String()
	return insert(exam, NewDynamoDbRepo().TableName)

}
