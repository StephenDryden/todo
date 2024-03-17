package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/validator"
	"github.com/stephendryden/todo/db"
	"github.com/stephendryden/todo/todo"
)

var validate *validator.Validate = validator.New()

func GetItem(table db.Table, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id, ok := request.PathParameters["id"]

	if ok {
		return GetTodo(id, table, request)
	} else {
		return GetTodos(table, request)
	}

}

func GetTodo(id string, table db.Table, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received GET todo request with id = %s", id)

	todo, err := table.GetTodo(id)
	if err != nil {
		return ServerError(err)
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

func GetTodos(table db.Table, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Received GET todos request")

	todos, err := table.GetTodos()
	if err != nil {
		return ServerError(err)
	}

	json, err := json.Marshal(todos)
	if err != nil {
		return ServerError(err)
	}
	log.Printf("Successfully fetched todos %s", json)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
	}, nil
}

func AddItem(table db.Table, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Print("Received PUT todo request")

	createTodo := todo.CreateTodo{}

	err := json.Unmarshal([]byte(request.Body), &createTodo)
	if err != nil {
		log.Printf("can't unmarshal body: %v", err)
		return ClientError(http.StatusUnprocessableEntity)
	}

	err = validate.Struct(&createTodo)
	if err != nil {
		log.Printf("invalid body: %v", err)
		return ClientError(http.StatusBadRequest)
	}
	log.Printf("received PUT request with item: %+v", &createTodo)

	response, err := table.AddTodo(createTodo)
	if err != nil {
		return ServerError(err)
	}

	log.Print("Successfully created todo")

	json, err := json.Marshal(response)
	if err != nil {
		return ServerError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(json),
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
