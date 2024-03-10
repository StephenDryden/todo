package stubs

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func StubGetTodo(
	tableName string, key map[string]types.AttributeValue, name string, description string, done string,
	raiseErr *testtools.StubError) testtools.Stub {
	return testtools.Stub{
		OperationName: "GetItem",
		Input:         &dynamodb.GetItemInput{TableName: aws.String(tableName), Key: key},
		Output: &dynamodb.GetItemOutput{Item: map[string]types.AttributeValue{
			"name":        &types.AttributeValueMemberS{Value: name},
			"description": &types.AttributeValueMemberN{Value: description},
			"done":        &types.AttributeValueMemberN{Value: done},
		},
		},
		Error: raiseErr,
	}
}
