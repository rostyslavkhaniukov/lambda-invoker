package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

type InvokerEvent struct {
	LambdaName string `json:"lambda_name"`
	Arguments string `json:"arguments"`
}

func HandleRequest(ctx context.Context, event InvokerEvent) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambdaService.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	result, err := client.Invoke(&lambdaService.InvokeInput{
		FunctionName: aws.String(event.LambdaName),
		Payload: []byte(event.Arguments),
	})
	if err != nil {
		return fmt.Sprintf("Error calling %s. Error: %s", event.LambdaName, err.Error()), nil
	}

	return fmt.Sprintf("Done! %s\n", string(result.Payload)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
