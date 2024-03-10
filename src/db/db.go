package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stephendryden/todo/todo"
)

type Table struct {
	DynamoDbClient *dynamodb.Client
	Name           string
}

func (table Table) GetTodo(id string) (todo.Todo, error) {
	todo := todo.Todo{
		Id: id,
	}

	response, err := table.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: todo.GetKey(), TableName: aws.String(table.Name),
	})
	if err != nil {
		log.Printf("Couldn't get todo with id %v. Here's why: %v\n", id, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &todo)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return todo, err
}

func (table Table) GetTodos() ([]todo.Todo, error) {
	todos := []todo.Todo{}

	response, err := table.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: &table.Name})

	if err != nil {
		log.Printf("Couldn't get todos. Here's why: %v\n", err)
	} else {
		for _, item := range response.Items {
			todo := todo.Todo{}
			err = attributevalue.UnmarshalMap(item, &todo)
			if err != nil {
				log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
			} else {
				todos = append(todos, todo)
			}
		}
	}
	return todos, err
}
