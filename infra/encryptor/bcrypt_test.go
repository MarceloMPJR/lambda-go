package encryptor_test

import (
	"testing"

	"github.com/MarceloMPJR/lambda-go/infra/encryptor"
	"golang.org/x/crypto/bcrypt"
)

func TestBCrypt(t *testing.T) {
	tests := []struct {
		givenCost     int
		givenPassword string
		expectedErr   error
	}{
		{1, "password_1", nil},
		{2, "password_2", nil},
		{3, "password_3", nil},
		{4, "password_4", nil},
		{5, "password_5", nil},
		{6, "password_6", nil},
		{7, "password_7", nil},
		{8, "password_8", nil},
		{9, "password_9", nil},
		{10, "password_10", nil},
		{50, "password_50", bcrypt.InvalidCostError(50)},
	}

	for _, test := range tests {
		bcrypt := encryptor.NewBCrypt(test.givenCost)

		result1, err := bcrypt.Encrypt(test.givenPassword)

		if test.expectedErr != nil && err == nil {
			t.Fatal("expected error, but not occurres")
		}

		if test.expectedErr == nil && err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		result2, err := bcrypt.Encrypt(test.givenPassword)

		if test.expectedErr != nil {
			if err == nil {
				t.Fatal("expected error, but not occurres")
			}

			if test.expectedErr != err {
				t.Errorf("result error: %v, expected error: %v", err, test.expectedErr)
			}
			continue
		}

		if test.expectedErr == nil && err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result1 == result2 {
			t.Errorf("result 1 and result 2 should be differents, but not are")
		}

		resultComp1, err := bcrypt.Compare(test.givenPassword, result1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !resultComp1 {
			t.Errorf("encrypt 1 generated a invalid hash")
		}

		resultComp2, err := bcrypt.Compare(test.givenPassword, result1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !resultComp2 {
			t.Errorf("encrypt 2 generated a invalid hash")
		}
	}
}
