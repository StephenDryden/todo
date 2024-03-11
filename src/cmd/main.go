package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stephendryden/todo/db"
	"github.com/stephendryden/todo/handlers"
)

func main() {

	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Received req %#v", request)

	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	config.LoadDefaultConfig(context.TODO())
	table := db.Table{
		DynamoDbClient: dynamodb.NewFromConfig(sdkConfig),
		Name:           "todo",
	}

	switch request.HTTPMethod {
	case "GET":
		return handlers.GetItem(table, request)
	default:
		return handlers.ClientError(http.StatusMethodNotAllowed)
	}
}
