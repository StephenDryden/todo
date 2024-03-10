package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stephendryden/todo/db"
	"github.com/stephendryden/todo/handlers"
)

func main() {

	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Received req %#v", request)

	sdkConfig := aws.Config{
		Region: "eu-west-1",
	}

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
