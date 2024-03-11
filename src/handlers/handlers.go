package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stephendryden/todo/db"
)

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
