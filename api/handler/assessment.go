package handler

import (
	"encoding/json"
	"net/http"

	"github.com/NicolasFkm/transaction-assessment-go/internal/assess"
	"github.com/NicolasFkm/transaction-assessment-go/internal/model"
	"github.com/NicolasFkm/transaction-assessment-go/internal/utils"
)

var logger = utils.GetLogger()

func AssessTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	transactionData := &model.AssessmentRequest{}
	err := json.NewDecoder(r.Body).Decode(transactionData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transactionHistory := map[int](model.UserTransaction){}
	assessmentResults := model.AssessmentResponse{}

	for _, transaction := range transactionData.Transactions {
		userTransaction, ok := transactionHistory[transaction.UserId]

		if ok == false {
			transactionHistory[transaction.UserId] = model.UserTransaction{
				TotalAmount: 0,
				CardsUsed:   []int{},
			}

			userTransaction = transactionHistory[transaction.UserId]
		}
		assessment, err := assess.AssessRisk(transaction, &userTransaction)
		transactionHistory[transaction.UserId] = userTransaction

		if err != nil {
			logger.Error("An error occurred")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		assessmentResults.RiskRatings = append(assessmentResults.RiskRatings, assessment)
	}

	json.NewEncoder(w).Encode(assessmentResults)
}
