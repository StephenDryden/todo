package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router)
}

type Todo struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
	Status      bool   `json:"status" dynamodbav:"status"`
}

type UpdateTodo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      bool   `json:"status" validate:"required"`
}

type CreateTodo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
