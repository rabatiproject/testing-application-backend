package implementation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"log"
)

type dynamoExamQuestionRepo struct {
	TableName string
}

type examQuestionDataType struct {
	Id         string `json:"id"`
	QuestionId string `json:"question_id"`
	ExamId     string `json:"exam_id"`
}

func NewExamQuestionRepo() *dynamoExamQuestionRepo {
	return &dynamoExamQuestionRepo{
		TableName: "EXAM_QUESTIONS",
	}
}
func (q *dynamoExamQuestionRepo) AddToExam(questionId, examId string) error {
	dataType := &examQuestionDataType{uuid.New().String(), examId, questionId}
	return insert(dataType, NewExamQuestionRepo().TableName)
}

func (d *dynamoExamQuestionRepo) ExamContainsQuestions(examId string, questionId string) bool {
	dynamodbClient := getNewClient()

	filt := expression.Name("exam_id").Equal(expression.Value(examId))

	proj := expression.NamesList(expression.Name("question_id"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(NewExamQuestionRepo().TableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamodbClient.Scan(params)

	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	for _, item := range result.Items {
		if *item["question_id"].S == questionId {
			return true
		}
	}

	return false
}
