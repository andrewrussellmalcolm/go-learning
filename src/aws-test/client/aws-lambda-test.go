package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// Request : request for service from handler
type Request struct {
	Message string
}

// Response : response to request
type Response struct {
	Message string
}

// AWSLambda :
type AWSLambda struct {
	lambda *lambda.Lambda
}

func (awsLambda *AWSLambda) init() {

	// Create Lambda service client
	session := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))

	awsLambda.lambda = lambda.New(session)

}

func (awsLambda *AWSLambda) invoke(request Request, response *Response) {

	payload, err := json.Marshal(request)

	if err != nil {
		fmt.Println("Error marshalling aws-lambda-test request")
		os.Exit(0)
	}

	result, err := awsLambda.lambda.Invoke(&lambda.InvokeInput{FunctionName: aws.String("aws-lambda-test"), Payload: payload})

	if err != nil {
		fmt.Println("Error calling aws-lambda-test")
		os.Exit(0)
	}

	err = json.Unmarshal(result.Payload, &response)

	if err != nil {
		fmt.Println("Error unmarshalling aws-lambda-test response")
		os.Exit(0)
	}
}

func main() {

	request := Request{Message: "Hello Andrew"}
	response := Response{}
	awsLambda := AWSLambda{}

	awsLambda.init()
	awsLambda.invoke(request, &response)
	fmt.Println(response.Message)
}
