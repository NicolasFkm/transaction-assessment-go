package validator

import (
	"testing"

	"github.com/NicolasFkm/transaction-assessment-go/internal/validator"
)

func Test_CheckAge(t *testing.T) {
	tests := []struct {
		testName       string
		age            int
		expectedResult bool
		errorMessage   string
	}{
		{
			"should return false when is not an adult",
			10,
			false,
			"",
		},
	}

	for _, testCase := range tests {
		result := validator.CheckIsAdult(testCase.age)
		if result != testCase.expectedResult {
			t.Errorf(testCase.testName, "wanted", testCase.expectedResult, "received", result)
		}
	}
}
