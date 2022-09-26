package token

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
