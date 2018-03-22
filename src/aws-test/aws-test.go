package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Message string
}

type Response struct {
	Message string
}

func Handler(request Request) (Response, error) {

	log.Printf("Processing Lambda request %s\n", request.Message)

	return Response{Message: request.Message}, nil
}

func main() {
	lambda.Start(Handler)
}
