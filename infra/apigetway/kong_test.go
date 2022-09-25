package apigetway_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarceloMPJR/lambda-go/infra/apigetway"
)

func TestKong_GetConsumerInfo(t *testing.T) {
	t.Run("when username is valid", func(t *testing.T) {
		username := "test"

		t.Run("when type of auth is valid", func(t *testing.T) {
			authType := "jwt"

			t.Run("when server response success", func(t *testing.T) {
				expectedKey := "KeY"
				srvTest := serverTestHTTP(username, authType, expectedKey, http.StatusOK)

				apiGetway := apigetway.NewKong(srvTest.URL)
				out := apiGetway.GetConsumerInfo(apigetway.ConsumerInfoInput{UserName: username, AuthType: authType})

				if out.Error != nil {
					t.Fatalf("unexpected error: %v", out.Error)
				}

				if out.Key != expectedKey {
					t.Errorf("result: %s, expected: %s", out.Key, expectedKey)
				}
			})

			t.Run("when server response error", func(t *testing.T) {
				srvTest := serverTestHTTP(username, authType, "", http.StatusInternalServerError)
				apiGetway := apigetway.NewKong(srvTest.URL)
				out := apiGetway.GetConsumerInfo(apigetway.ConsumerInfoInput{UserName: username, AuthType: authType})

				checkError(t, apigetway.KongCommunicationFailure, out)
			})
		})

		t.Run("when type of auth is invalid", func(t *testing.T) {
			authType := "any"

			srvTest := serverTestHTTP(username, authType, "", http.StatusInternalServerError)
			apiGetway := apigetway.NewKong(srvTest.URL)
			out := apiGetway.GetConsumerInfo(apigetway.ConsumerInfoInput{UserName: username, AuthType: authType})

			checkError(t, apigetway.KongCommunicationFailure, out)
		})
	})

	t.Run("when username is invalid", func(t *testing.T) {
		username := "any"

		t.Run("when type of auth is valid", func(t *testing.T) {
			authType := "jwt"

			srvTest := serverTestHTTP(username, authType, "", http.StatusInternalServerError)
			apiGetway := apigetway.NewKong(srvTest.URL)
			out := apiGetway.GetConsumerInfo(apigetway.ConsumerInfoInput{UserName: username, AuthType: authType})

			checkError(t, apigetway.KongCommunicationFailure, out)
		})

		t.Run("when type of auth is invalid", func(t *testing.T) {
			authType := "any"

			srvTest := serverTestHTTP(username, authType, "", http.StatusInternalServerError)
			apiGetway := apigetway.NewKong(srvTest.URL)
			out := apiGetway.GetConsumerInfo(apigetway.ConsumerInfoInput{UserName: username, AuthType: authType})

			checkError(t, apigetway.KongCommunicationFailure, out)
		})
	})
}

func checkError(t *testing.T, expectedError string, result apigetway.ConsumerInfoOutput) {
	t.Helper()

	if result.Error == nil {
		t.Fatal("expected error, but not occurres")
	}

	if result.Error.Error() != expectedError {
		t.Errorf("error returned: %s, expected: %s", result.Error, expectedError)
	}
}

func serverTestHTTP(userName, authType, key string, status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if status != http.StatusOK {
			w.WriteHeader(status)
			fmt.Fprint(w, "error")
		}

		if userName == "test" && authType == "jwt" {
			responseBody := fmt.Sprintf(`{"key": "%s"}`, key)
			fmt.Fprint(w, responseBody)

			return
		}

		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "not found")
	}))
}
