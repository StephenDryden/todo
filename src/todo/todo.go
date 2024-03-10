package todo

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Todo struct {
	Id          string `dynamodbav:"id"`
	Name        string `dynamodbav:"name"`
	Description string `dynamodbav:"description"`
	Done        string `dynamodbav:"done"`
}

func (todo Todo) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(todo.Id)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id}
}
