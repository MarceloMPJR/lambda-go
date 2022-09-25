package app

import "github.com/MarceloMPJR/lambda-go/infra/token"

var CurrentApp = App{}

type App struct {
	TokenGenerator token.TokenGenerator
}
