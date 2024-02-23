package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/stephendryden/todo/handlers"
)

func main() {
	lambda.Start(handlers.Router)
}
