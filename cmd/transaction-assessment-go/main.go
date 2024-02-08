package main

import (
	"net/http"

	"github.com/NicolasFkm/transaction-assessment-go/api/handler"
)

func main() {
	http.HandleFunc("/assessment", handler.AssessTransactions)
	http.ListenAndServe(":8080", nil)
}
