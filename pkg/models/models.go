package models

type TransactionRequest struct {
	ExternalId      string `json:"id,omitempty"`
	Amount          int64  `json:"amount"`
	Description     string `json:"description"`
	CurrencyCode    int32  `json:"currencyCode,omitempty"`
	Balance         int64  `json:"balance,omitempty"`
	MCC             int32  `json:"mcc,omitempty"`
	TransactionTime int64  `json:"time,omitempty"`
}

const MONOBANK_SOURCE = "monobank"
const PRIVATBANK_SOURCE = "privatbank"

type Client struct {
	Source     string     // "monobank", "privatbank"
	ExternalId string     `json:"clientId,omitempty"`
	Name       string     `json:"name,omitempty"`
	Accounts   []*Account `json:"accounts,omitempty"`
	Jars       []*Jar     `json:"jars,omitempty"`
}
type Account struct {
	Source       string   // "monobank", "privatbank"
	ExternalId   string   `json:"id,omitempty"`
	CurrencyCode int32    `json:"currencyCode,omitempty"` // 980 (UAH), 840 (USD)
	Balance      int64    `json:"balance,omitempty"`      // in kopiykas/cents
	MaskedPan    []string `json:"maskedPan,omitempty"`    // last 4 card's digits
	Name         string   `json:"type,omitempty"`
}

type Jar struct {
	ExternalId   string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	CurrencyCode int32  `json:"currencyCode,omitempty"`
	Balance      int64  `json:"balance,omitempty"`
	Goal         int64  `json:"goal"`
}
