package implementation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"log"
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

func (d *dynamoDbExamRepo) ExamExists(examId string) bool {
	dynamodbClient := getNewClient()

	filt := expression.Name("id").Equal(expression.Value(examId))

	proj := expression.NamesList(expression.Name("id"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(NewDynamoDbRepo().TableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamodbClient.Scan(params)

	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	return len(result.Items) > 0
}
