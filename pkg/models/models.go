package models

type TransactionRequest struct {
	ExternalId      string `json:"external_id"`
	Amount          int64  `json:"amount"`
	Description     string `json:"description"`
	CurrencyCode    int32  `json:"currency_code"`
	Balance         int64  `json:"balance"`
	MCC             int32  `json:"mcc"`
	TransactionTime int64  `json:"time"`
}

const MONOBANK_SOURCE = "monobank"
const PRIVATBANK_SOURCE = "privatbank"
