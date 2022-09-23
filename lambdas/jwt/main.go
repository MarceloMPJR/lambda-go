package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtAlgo = "HS256"
	jwtType = "JWT"

	jwtKey = "ar9rVhd9ORBBJu1T9Eon2SLlCovhtO57" // BY USER
)

func main() {
	lambda.Start(handlerGenerate)
}

func handlerGenerate(ctx context.Context) (string, error) {
	claims := &jwt.MapClaims{
		"name": "testes",
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["alg"] = jwtAlgo
	token.Header["kid"] = jwtKey
	token.Header["typ"] = jwtType

	return token.SignedString([]byte(hmacSecret()))
}

func hmacSecret() string {
	return os.Getenv("HMAC_SECRET")
}
