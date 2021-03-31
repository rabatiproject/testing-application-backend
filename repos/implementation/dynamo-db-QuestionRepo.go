package implementation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"log"
)

type questionRepo struct {
	TableName string
}

func NewQuestionRepo() *questionRepo {
	return &questionRepo{
		TableName: "QUESTIONS",
	}
}

func (q *questionRepo) CreateMCQuestion(question *base.MultipleChoiceQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionMultipleChoice
	return insert(question, NewQuestionRepo().TableName)
}

func (q *questionRepo) CreateOEQuestion(question *base.OpenEndedQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionOpenEnded
	return insert(question, NewQuestionRepo().TableName)
}

func (q *questionRepo) CreatePQuestion(question *base.ProgrammingQuestion) error {
	question.Id = uuid.New().String()
	question.Type = base.QuestionProgramming
	return insert(question, NewQuestionRepo().TableName)
}

func (q *questionRepo) QuestionExists(questionId string) bool {
	dynamodbClient := getNewClient()

	filt := expression.Name("id").Equal(expression.Value(questionId))

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
		TableName:                 aws.String(NewQuestionRepo().TableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamodbClient.Scan(params)

	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	return len(result.Items) > 0
}

func (q *questionRepo) AddToExam(questionId, examId string) error {
	panic("implement me")
}
