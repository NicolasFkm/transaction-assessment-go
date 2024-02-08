package assess

import "github.com/NicolasFkm/transaction-assessment-go/internal/model"

func AssessRisk(transaction model.Transaction, userTransaction *model.UserTransaction) (string, error) {
	amount_us_dollars := float64(transaction.AmountInCents / 100)

	userTransaction.TotalAmount += amount_us_dollars
	userTransaction.CardsUsed = append(userTransaction.CardsUsed, transaction.CardId)

	cardsUsed := make(map[int]bool)
	for card := range userTransaction.CardsUsed {
		cardsUsed[card] = true
	}

	if len(cardsUsed) > 2 {
		return "high", nil
	}
	if userTransaction.TotalAmount > 20000 {
		return "high", nil
	}

	if amount_us_dollars > 10000 {
		return "high", nil
	}

	if amount_us_dollars > 5000 {
		return "medium", nil
	}

	if len(cardsUsed) > 1 {
		return "medium", nil
	}
	if userTransaction.TotalAmount > 10000 {
		return "medium", nil
	}

	return "low", nil

}
