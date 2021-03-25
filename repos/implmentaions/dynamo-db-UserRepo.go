package implmentaions

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/rabatiproject/testing-application-backend/model/base"
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
	dynamodbClient := getNewClient()
	user.Id = uuid.New().String()
	attributeValue, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		fmt.Println(err)
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(NewUserRepo().TableName),
		Item:      attributeValue,
	}

	_, err2 := dynamodbClient.PutItem(input)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	return nil
}
