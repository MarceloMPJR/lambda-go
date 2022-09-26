package users

import (
	"github.com/MarceloMPJR/lambda-go/entity"
	"github.com/MarceloMPJR/lambda-go/infra/apigetway"
	"github.com/MarceloMPJR/lambda-go/infra/token"
)

type AuthorizeOutput struct {
	Token string
	Error error
}

type AuthorizeService struct {
	tokenGenerator token.TokenGenerator
	consumerInfo   apigetway.ConsumerInfo
}

type Authorizer interface {
	Authorize(entity.User) AuthorizeOutput
}
