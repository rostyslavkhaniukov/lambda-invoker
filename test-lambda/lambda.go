package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, name string) (string, error) {
	fmt.Printf("Hello %s!", name)
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
