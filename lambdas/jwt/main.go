package main

import (
	"context"

	"github.com/MarceloMPJR/lambda-go/entity"
	"github.com/MarceloMPJR/lambda-go/infra/app"
	"github.com/MarceloMPJR/lambda-go/services/users"
	"github.com/aws/aws-lambda-go/lambda"
)

type UserParams struct {
	Name string `json:"name"`
}

func handlerGenerate(ctx context.Context, userParam UserParams) (string, error) {
	authService := users.NewAuthorizeService(app.CurrentApp.TokenGenerator, app.CurrentApp.ConsumerInfo)
	user := entity.User{Name: userParam.Name}
	out := authService.Authorize(user)

	return out.Token, out.Error
}

func main() {
	app.SetupTokenGenerator()
	app.SetupConsumerInfo()

	lambda.Start(handlerGenerate)
}
