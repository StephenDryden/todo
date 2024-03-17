package todo

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Todo struct {
	Id          string `dynamodbav:"id"`
	Name        string `dynamodbav:"name"`
	Description string `dynamodbav:"description"`
	Done        string `dynamodbav:"done"`
}

type UpdateTodo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Done        bool   `json:"done" validate:"required"`
}

type CreateTodo struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (todo Todo) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(todo.Id)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id}
}

func (createTodo CreateTodo) NewId() string {
	return uuid.NewString()
}
