package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stephendryden/todo/todo"
)

// GetItem returns a todo item from dynamodb using a given id
func GetItem(ctx context.Context, id string) (*todo.Todo, error) {
	key, err := attributevalue.Marshal(id)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": key,
		},
	}

	log.Printf("Calling Dynamodb with input: %v", input)
	result, err := db.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}
	log.Printf("Executed GetItem DynamoDb successfully. Result: %#v", result)

	if result.Item == nil {
		return nil, nil
	}

	todo := new(todo.Todo)
	err = attributevalue.UnmarshalMap(result.Item, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
