package token

type Payload map[string]interface{}

type TokenGeneratorInput struct {
	Key     string
	Payload Payload
}

type TokenGeneratorOutput struct {
	Token string
	Error error
}

type TokenGenerator interface {
	Generate(TokenGeneratorInput) TokenGeneratorOutput
}
