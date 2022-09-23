package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlerHello)
}

func handlerHello(ctx context.Context) (string, error) {
	return "Hello", nil
}
