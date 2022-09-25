package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenGeneratorInput struct {
	Key     string
	Payload interface{}
}

type TokenGeneratorOutput struct {
	Token string
	Error error
}

type TokenGenerator interface {
	Generate(TokenGeneratorInput) TokenGeneratorOutput
}

const (
	jwtAlgo = "HS256"
	jwtType = "JWT"
)

type JWT struct {
	secret []byte
}

func NewJWT(secret []byte) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) Generate(in TokenGeneratorInput) (out TokenGeneratorOutput) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, in.Payload.(*jwt.MapClaims))

	token.Header["kid"] = in.Key

	token.Header["typ"] = jwtType
	token.Header["alg"] = jwtAlgo

	out.Token, out.Error = token.SignedString(j.secret)

	return
}
