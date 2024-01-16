package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo/db"
	"todo/todo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func ProcessGet(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := req.PathParameters["id"]
	if !ok {
		return ProcessGetTodos(ctx)
	} else {
		return ProcessGetTodo(ctx, id)
	}
}

func ProcessGetTodo(ctx context.Context, id string) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received GET todo request with id = %s", id)

	todo, err := db.GetItem(ctx, id)
	if err != nil {
		return ServerError(err)
	}

	if todo == nil {
		return ClientError(http.StatusNotFound)
	}

	json, err := json.Marshal(todo)
	if err != nil {
		return ServerError(err)
	}
	log.Printf("Successfully fetched todo item %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func ProcessGetTodos(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	log.Print("Received GET todos request")

	todos, err := db.ListItems(ctx)
	if err != nil {
		return ServerError(err)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		return ServerError(err)
	}
	log.Printf("Successfully fetched todos: %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func ProcessPost(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var createTodo todo.CreateTodo
	err := json.Unmarshal([]byte(req.Body), &createTodo)
	if err != nil {
		log.Printf("Can't unmarshal body: %v", err)
		return ClientError(http.StatusUnprocessableEntity)
	}

	err = validate.Struct(&createTodo)
	if err != nil {
		log.Printf("Invalid body: %v", err)
		return ClientError(http.StatusBadRequest)
	}
	log.Printf("Received POST request with item: %+v", createTodo)

	res, err := db.InsertItem(ctx, createTodo)
	if err != nil {
		return ServerError(err)
	}
	log.Printf("Inserted new todo: %+v", res)

	json, err := json.Marshal(res)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(json),
		Headers: map[string]string{
			"Location": fmt.Sprintf("/todo/%s", res.Id),
		},
	}, nil
}

func ProcessDelete(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := req.PathParameters["id"]
	if !ok {
		return ClientError(http.StatusBadRequest)
	}
	log.Printf("Received DELETE request with id = %s", id)

	todo, err := db.DeleteItem(ctx, id)
	if err != nil {
		return ServerError(err)
	}

	if todo == nil {
		return ClientError(http.StatusNotFound)
	}

	json, err := json.Marshal(todo)
	if err != nil {
		return ServerError(err)
	}
	log.Printf("Successfully deleted todo item %+v", todo)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func ProcessPut(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := req.PathParameters["id"]
	if !ok {
		return ClientError(http.StatusBadRequest)
	}

	var updateTodo todo.UpdateTodo
	err := json.Unmarshal([]byte(req.Body), &updateTodo)
	if err != nil {
		log.Printf("Can't unmarshal body: %v", err)
		return ClientError(http.StatusUnprocessableEntity)
	}

	err = validate.Struct(&updateTodo)
	if err != nil {
		log.Printf("Invalid body: %v", err)
		return ClientError(http.StatusBadRequest)
	}
	log.Printf("Received PUT request with item: %+v", updateTodo)

	res, err := db.UpdateItem(ctx, id, updateTodo)
	if err != nil {
		return ServerError(err)
	}

	if res == nil {
		return ClientError(http.StatusNotFound)
	}

	log.Printf("Updated todo: %+v", res)

	json, err := json.Marshal(res)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
		Headers: map[string]string{
			"Location": fmt.Sprintf("/todo/%s", res.Id),
		},
	}, nil
}
func ClientError(status int) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(status),
		StatusCode: status,
	}, nil
}

func ServerError(err error) (events.APIGatewayProxyResponse, error) {
	log.Println(err.Error())

	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusInternalServerError),
		StatusCode: http.StatusInternalServerError,
	}, nil
}
