package db

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
)

func deleteItem(ctx context.Context, id string) (*Todo, error) {
	key, err := attributevalue.Marshal(id)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(TableName),
		Key: map[string]types.AttributeValue{
			"id": key,
		},
		ReturnValues: types.ReturnValue(*aws.String("ALL_OLD")),
	}

	res, err := db.DeleteItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if res.Attributes == nil {
		return nil, nil
	}

	todo := new(Todo)
	err = attributevalue.UnmarshalMap(res.Attributes, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func updateItem(ctx context.Context, id string, updateTodo UpdateTodo) (*Todo, error) {
	key, err := attributevalue.Marshal(id)
	if err != nil {
		return nil, err
	}

	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(
			expression.Name("name"),
			expression.Value(updateTodo.Name),
		).Set(
			expression.Name("description"),
			expression.Value(updateTodo.Description),
		).Set(
			expression.Name("status"),
			expression.Value(updateTodo.Status),
		),
	).WithCondition(
		expression.Equal(
			expression.Name("id"),
			expression.Value(id),
		),
	).Build()
	if err != nil {
		return nil, err
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]types.AttributeValue{
			"id": key,
		},
		TableName:                 aws.String(TableName),
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ConditionExpression:       expr.Condition(),
		ReturnValues:              types.ReturnValue(*aws.String("ALL_NEW")),
	}

	res, err := db.UpdateItem(ctx, input)
	if err != nil {
		var smErr *smithy.OperationError
		if errors.As(err, &smErr) {
			var condCheckFailed *types.ConditionalCheckFailedException
			if errors.As(err, &condCheckFailed) {
				return nil, nil
			}
		}

		return nil, err
	}

	if res.Attributes == nil {
		return nil, nil
	}

	todo := new(Todo)
	err = attributevalue.UnmarshalMap(res.Attributes, todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
