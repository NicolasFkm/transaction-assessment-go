package assess

import (
	"testing"

	"github.com/NicolasFkm/transaction-assessment-go/internal/assess"
	"github.com/NicolasFkm/transaction-assessment-go/internal/model"
)

func Test_AssessTransaction(t *testing.T) {
	tests := []struct {
		testName        string
		transaction     model.Transaction
		userTransaction model.UserTransaction
		expectedResult  string
	}{
		{
			"should return false when is not an adult",
			model.Transaction{Id: 1, UserId: 1, AmountInCents: 200000, CardId: 1},
			model.UserTransaction{TotalAmount: 0, CardsUsed: []int{}},
			"low",
		},
	}

	for _, testCase := range tests {
		result, err := assess.AssessRisk(testCase.transaction, &testCase.userTransaction)
		if result != testCase.expectedResult || err != nil {
			t.Errorf(testCase.testName, "wanted", testCase.expectedResult, "received", result)
		}
	}
}
