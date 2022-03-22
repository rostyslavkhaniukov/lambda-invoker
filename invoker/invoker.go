package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambdaService.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	result, err := client.Invoke(&lambdaService.InvokeInput{
		FunctionName: aws.String("test-lambda-web"),
	})
	if err != nil {
		return fmt.Sprintf("Error calling test-lambda-web. Error: %s", err.Error()), nil
	}

	return fmt.Sprintf("All OK! %s\n", string(result.Payload)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
