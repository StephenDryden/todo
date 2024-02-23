package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/stephendryden/todo/todo"
)

// PostItem adds a new todo item to dynamodb
// that newly added item is returned
func PostItem(ctx context.Context, createTodo todo.CreateTodo) (*todo.Todo, error) {
	todo := todo.Todo{
		Name:        createTodo.Name,
		Description: createTodo.Description,
		Status:      false,
		Id:          uuid.NewString(),
	}

	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item:      item,
	}

	res, err := db.PutItem(ctx, input)
	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalMap(res.Attributes, &todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
