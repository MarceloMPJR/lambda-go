package users

import (
	"errors"
	"time"

	"github.com/MarceloMPJR/lambda-go/entity"
	"github.com/MarceloMPJR/lambda-go/infra/apigetway"
	"github.com/MarceloMPJR/lambda-go/infra/token"
)

type UserPayload map[string]interface{}

const (
	InvalidUserMessage = "username/password invalid"

	authType = "jwt"

	validUser     = "testes"
	validPassword = "123456"
)

func NewAuthorizeService(tokenGenerator token.TokenGenerator, consumerInfo apigetway.ConsumerInfo) *AuthorizeService {
	return &AuthorizeService{tokenGenerator: tokenGenerator, consumerInfo: consumerInfo}
}

func (a *AuthorizeService) Authorize(user entity.User) (out AuthorizeOutput) {
	if user.Name == validUser && user.Password == validPassword {
		key, err := a.keyByUsername(user.Name)
		if err != nil {
			out.Error = err
			return
		}

		t := a.tokenGenerator.Generate(token.TokenGeneratorInput{Key: key, Payload: a.buildPayload(user)})
		return AuthorizeOutput(t)
	}

	out.Error = errors.New(InvalidUserMessage)
	return
}

func (a *AuthorizeService) keyByUsername(username string) (string, error) {
	consumerInfoInput := apigetway.ConsumerInfoInput{UserName: username, AuthType: authType}
	consuerInfoOutput := a.consumerInfo.GetConsumerInfo(consumerInfoInput)

	return consuerInfoOutput.Key, consuerInfoOutput.Error
}

func (AuthorizeService) buildPayload(user entity.User) token.Payload {
	return token.Payload{
		"name": user.Name,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}
}
