package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s. \n", request.RequestContext.RequestID)
	fmt.Println("Body = ", request.Body)

	fmt.Println("\n Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s - ", key, value)
	}

	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}