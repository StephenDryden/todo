package db

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
	"github.com/stephendryden/todo/stubs"
	"github.com/stephendryden/todo/todo"
)

func enterTest() (*testtools.AwsmStubber, *Table) {
	stubber := testtools.NewStubber()
	table := &Table{Name: "test-table", DynamoDbClient: dynamodb.NewFromConfig(*stubber.SdkConfig)}
	return stubber, table
}

func TestTable_GetTodo(t *testing.T) {
	t.Run("NoErrors", func(t *testing.T) { GetTodo(nil, t) })
	t.Run("TestError", func(t *testing.T) { GetTodo(&testtools.StubError{Err: errors.New("TestError")}, t) })
}

func GetTodo(raiseErr *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	todo := todo.Todo{
		Id:          "1",
		Name:        "First Todo",
		Description: "This is the first todo",
		Done:        "false",
	}

	stubber.Add(stubs.StubGetTodo(table.Name, todo.GetKey(), todo.Name, todo.Description,
		todo.Done, raiseErr))

	gotTodo, err := table.GetTodo(todo.Id)

	testtools.VerifyError(err, raiseErr, t)
	if err == nil {
		if gotTodo.Name != todo.Name || gotTodo.Description != todo.Description {
			t.Errorf("got %s but expected %s", gotTodo, todo)
		}
	}

	testtools.ExitTest(stubber, t)
}

// Todo: fix tests
func TestTableBasics_AddTodo(t *testing.T) {
	t.Run("NoErrors", func(t *testing.T) { AddTodo(nil, t) })
	t.Run("TestError", func(t *testing.T) { AddTodo(&testtools.StubError{Err: errors.New("TestError")}, t) })
}

func AddTodo(raiseErr *testtools.StubError, t *testing.T) {
	stubber, table := enterTest()

	todo := todo.CreateTodo{Name: "test todo", Description: "this is a test todo"}

	item, marshErr := attributevalue.MarshalMap(todo)
	if marshErr != nil {
		panic(marshErr)
	}

	stubber.Add(stubs.StubAddTodo(table.Name, item, raiseErr))

	_, err := table.AddTodo(todo)

	testtools.VerifyError(err, raiseErr, t)
	testtools.ExitTest(stubber, t)
}
