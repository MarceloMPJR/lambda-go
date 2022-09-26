package users_test

import (
	"errors"
	"testing"

	"github.com/MarceloMPJR/lambda-go/entity"
	"github.com/MarceloMPJR/lambda-go/infra/apigetway"
	"github.com/MarceloMPJR/lambda-go/infra/token"
	"github.com/MarceloMPJR/lambda-go/services/users"
)

func Test_Authorize(t *testing.T) {
	t.Run("when user is authorized", func(t *testing.T) {
		user := entity.User{Name: "test", Password: "123456"}

		t.Run("when get info consumer has success", func(t *testing.T) {
			mockConsumerInfo := &MockConsumerInfo{KeyResult: "KeY", ErrorResult: nil}

			t.Run("when token generator has success", func(t *testing.T) {
				expectedToken := "ToKeN"
				mockTokenGenerator := &MockTokenGenerator{TokenResult: expectedToken, ErrorResult: nil}

				userService := users.NewAuthorizeService(mockTokenGenerator, mockConsumerInfo)
				out := userService.Authorize(user)

				if out.Error != nil {
					t.Fatalf("unexpected error: %v", out.Error)
				}

				if out.Token != expectedToken {
					t.Errorf("result: %s, expected: %s", out.Token, expectedToken)
				}
			})

			t.Run("when token generator has error", func(t *testing.T) {
				expectedError := "error generator"
				mockTokenGenerator := &MockTokenGenerator{TokenResult: "", ErrorResult: errors.New(expectedError)}

				userService := users.NewAuthorizeService(mockTokenGenerator, mockConsumerInfo)
				out := userService.Authorize(user)

				checkError(t, expectedError, out)
			})
		})

		t.Run("when get info consumer has error", func(t *testing.T) {
			expectedError := "error consumer"
			mockConsumerInfo := &MockConsumerInfo{ErrorResult: errors.New(expectedError)}

			t.Run("when token generator has success", func(t *testing.T) {
				expectedToken := "ToKeN"
				mockTokenGenerator := &MockTokenGenerator{TokenResult: expectedToken, ErrorResult: nil}

				userService := users.NewAuthorizeService(mockTokenGenerator, mockConsumerInfo)
				out := userService.Authorize(user)

				checkError(t, expectedError, out)
			})

			t.Run("when token generator has error", func(t *testing.T) {
				mockTokenGenerator := &MockTokenGenerator{ErrorResult: errors.New("error generator")}

				userService := users.NewAuthorizeService(mockTokenGenerator, mockConsumerInfo)
				out := userService.Authorize(user)

				checkError(t, expectedError, out)
			})
		})
	})

	t.Run("when user isn't authorized", func(t *testing.T) {
		user := entity.User{Name: "invalid", Password: "invalid"}

		userService := users.NewAuthorizeService(nil, nil)
		out := userService.Authorize(user)

		checkError(t, users.InvalidUserMessage, out)
	})
}

func checkError(t *testing.T, expectedError string, out users.AuthorizeOutput) {
	t.Helper()

	if out.Error == nil {
		t.Fatal("expected error, but not occurres")
	}

	if out.Error.Error() != expectedError {
		t.Errorf("result: %s, expected: %s", out.Error.Error(), expectedError)
	}
}

type MockConsumerInfo struct {
	KeyResult   string
	ErrorResult error
}

func (mc *MockConsumerInfo) GetConsumerInfo(apigetway.ConsumerInfoInput) (out apigetway.ConsumerInfoOutput) {
	out.Key = mc.KeyResult
	out.Error = mc.ErrorResult

	return
}

type MockTokenGenerator struct {
	TokenResult string
	ErrorResult error
}

func (mt *MockTokenGenerator) Generate(token.TokenGeneratorInput) (out token.TokenGeneratorOutput) {
	out.Token = mt.TokenResult
	out.Error = mt.ErrorResult

	return
}
