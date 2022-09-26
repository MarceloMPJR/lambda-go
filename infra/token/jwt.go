package token

import "github.com/golang-jwt/jwt/v4"

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
