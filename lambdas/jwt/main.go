package main

import (
	"context"

	"github.com/MarceloMPJR/lambda-go/entity"
	"github.com/MarceloMPJR/lambda-go/infra/app"
	"github.com/MarceloMPJR/lambda-go/services/users"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	User UserParams `json:"request_body_args"`
}

type UserParams struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func handlerGenerate(ctx context.Context, event Event) (string, error) {
	authService := users.NewAuthorizeService(app.CurrentApp.TokenGenerator, app.CurrentApp.ConsumerInfo)
	user := entity.User{Name: event.User.Name, Password: event.User.Password}
	out := authService.Authorize(user)

	return out.Token, out.Error
}

func main() {
	app.SetupTokenGenerator()
	app.SetupConsumerInfo()

	lambda.Start(handlerGenerate)
}
