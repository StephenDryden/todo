package main

import (
	"context"
	"log"
	"net/http"
	"todo/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router)
}

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received req %#v", req)

	switch req.HTTPMethod {
	case "GET":
		return handlers.ProcessGet(ctx, req)
	case "POST":
		return handlers.ProcessPost(ctx, req)
	case "DELETE":
		return handlers.ProcessDelete(ctx, req)
	case "PUT":
		return handlers.ProcessPut(ctx, req)
	default:
		return handlers.ClientError(http.StatusMethodNotAllowed)
	}
}
