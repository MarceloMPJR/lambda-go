package app

import (
	"os"

	"github.com/MarceloMPJR/lambda-go/infra/token"
)

func SetupTokenGenerator() {
	CurrentApp.TokenGenerator = token.NewJWT([]byte(os.Getenv("HMAC_SECRET")))
}
