package model

type Transaction struct {
	Id            int `json:"id"`
	UserId        int `json:"user_id"`
	AmountInCents int `json:"amount_us_cents"`
	CardId        int `json:"card_id"`
}
type UserTransaction struct {
	TotalAmount float64 `json:"total_amount"`
	CardsUsed   []int   `json:"cards_used"`
}
type AssessmentRequest struct {
	Transactions []Transaction `json:"transactions"`
}

type AssessmentResponse struct {
	RiskRatings []string `json:"risk_ratings"`
}
