package token_test

import (
	"testing"

	"github.com/MarceloMPJR/lambda-go/infra/token"
	"github.com/golang-jwt/jwt/v4"
)

func TestJWT_Generate(t *testing.T) {
	secret := "TESTE_1"
	key := "TESTE_2"
	payload := &jwt.MapClaims{
		"name": "username",
		"exp":  123456789,
	}
	input := token.TokenGeneratorInput{
		Key:     key,
		Payload: payload,
	}

	expectedHeader := "eyJhbGciOiJIUzI1NiIsImtpZCI6IlRFU1RFXzIiLCJ0eXAiOiJKV1QifQ"
	expectedBody := "eyJleHAiOjEyMzQ1Njc4OSwibmFtZSI6InVzZXJuYW1lIn0.zwFlJ295Iyl8wlvH22Ia1w0lt51J8Qnp_FT2i-lE3d4"
	expectedToken := expectedHeader + "." + expectedBody

	j := token.NewJWT([]byte(secret))
	out := j.Generate(input)

	if out.Error != nil {
		t.Fatalf("unexpected error: %v", out.Error)
	}

	if expectedToken != out.Token {
		t.Errorf("result: %v, expected: %v", out.Token, expectedToken)
	}
}
