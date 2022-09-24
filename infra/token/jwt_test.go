package token_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/MarceloMPJR/lambda-go/infra/token"
	"github.com/golang-jwt/jwt/v4"
)

func TestJWT_Generate(t *testing.T) {
	j := token.NewJWT("HS256", []byte("fP3Gk5d3lyFmLKDQQKafl5Z20iF9R1fa"))
	str, _ := j.Generate("ar9rVhd9ORBBJu1T9Eon2SLlCovhtO57", &jwt.MapClaims{
		"name": "testes",
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	})

	fmt.Println(str)
}
