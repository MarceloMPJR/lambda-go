package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type TokenGenerator interface {
	Generate(key string, payload interface{}) (string, error)
}

const jwtType = "JWT"

type JWT struct {
	algo   string
	secret []byte
}

func NewJWT(algo string, secret []byte) *JWT {
	return &JWT{algo: algo, secret: secret}
}

func (j *JWT) Generate(key string, payload interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload.(*jwt.MapClaims))
	token.Header["typ"] = jwtType

	token.Header["alg"] = j.algo
	token.Header["kid"] = key

	return token.SignedString(j.secret)
}
