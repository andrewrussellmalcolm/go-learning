package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Request : request for service from handler
type Request struct {
	Message string
}

// Response : response to request
type Response struct {
	Message string
}

// Handler :
func Handler(request Request) (Response, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.Message)

	return Response{Message: request.Message}, nil
}

func main() {
	lambda.Start(Handler)
}
