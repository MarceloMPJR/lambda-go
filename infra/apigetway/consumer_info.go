package apigetway

type ConsumerInfoInput struct {
	UserName string
	AuthType string
}

type ConsumerInfoOutput struct {
	Key   string
	Error error
}

type ConsumerInfo interface {
	GetConsumerInfo(ConsumerInfoInput) ConsumerInfoOutput
}
