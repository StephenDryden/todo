package db

import (
	"context"
	"todo/todo"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// ListItems lists all todo items in dynamodb
func ListItems(ctx context.Context) ([]todo.Todo, error) {
	todos := make([]todo.Todo, 0)
	var token map[string]types.AttributeValue

	for {
		input := &dynamodb.ScanInput{
			TableName:         aws.String(TableName),
			ExclusiveStartKey: token,
		}

		result, err := db.Scan(ctx, input)
		if err != nil {
			return nil, err
		}

		var fetchedTodos []todo.Todo
		err = attributevalue.UnmarshalListOfMaps(result.Items, &fetchedTodos)
		if err != nil {
			return nil, err
		}

		todos = append(todos, fetchedTodos...)
		token = result.LastEvaluatedKey
		if token == nil {
			break
		}
	}

	return todos, nil
}
