package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

type InvokerEvent struct {
	LambdaName string `json:"lambda_name"`
	Arguments  string `json:"arguments"`
}

func HandleRequest(ctx context.Context, event InvokerEvent) (string, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := lambdaService.New(sess, &aws.Config{Region: aws.String("us-east-1")})

	/*output, err := client.ListFunctions(&lambdaService.ListFunctionsInput{
		FunctionVersion: aws.String("ALL"),
	})
	if err != nil {
		return fmt.Sprintf("ListFunctions. Error: %s", err.Error()), nil
	}

	var functions []string
	for _, elem := range output.Functions {
		functions = append(functions, *elem.FunctionName)
	}*/

	payload, err := json.Marshal([]byte(event.Arguments))
	if err != nil {
		return fmt.Sprintf("Invalid CLI arguments. Error: %s", err.Error()), nil
	}

	result, err := client.Invoke(&lambdaService.InvokeInput{
		FunctionName: aws.String(event.LambdaName),
		Payload:      payload,
	})
	if err != nil {
		return fmt.Sprintf("Error calling %s. Error: %s", event.LambdaName, err.Error()), nil
	}

	return fmt.Sprintf("Done! %s\n", string(result.Payload)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
