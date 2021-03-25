package implementation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
	"log"
)

type userRepo struct {
	TableName string
}

func NewUserRepo() *userRepo {
	return &userRepo{
		TableName: "USERS",
	}
}

func (userRepo *userRepo) CreateUser(user *base.User) error {
	user.Id = uuid.New().String()
	return insert(user, NewUserRepo().TableName)
}

func (userRepo *userRepo) UserExists(email string) bool {
	return getUserFrom(email) != nil
}

func (userRepo *userRepo) GetUserFrom(email string) *base.User {
	return getUserFrom(email)
}

func getUserFrom(email string) *base.User {
	dynamodbClient := getNewClient()

	filt := expression.Name("email").Equal(expression.Value(email))

	proj := expression.NamesList(expression.Name("id"),
		expression.Name("name"),
		expression.Name("surname"),
		expression.Name("password"),
		expression.Name("email"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(NewUserRepo().TableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynamodbClient.Scan(params)

	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	for _, i := range result.Items {

		user := base.User{}
		err = dynamodbattribute.UnmarshalMap(i, &user)
		if err != nil {
			log.Printf("Got error unmarshalling: %s\n", err)
		}
		return &user
	}

	return nil
}
