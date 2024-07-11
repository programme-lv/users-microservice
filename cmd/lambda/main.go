package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/programme-lv/users-microservice/internal/handlers"
)

func main() {
	lambda.Start(handlers.HandleRequest)
}
