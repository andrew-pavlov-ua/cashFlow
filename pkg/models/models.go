package models

import (
	"fmt"
	"time"
)

type MonoWebhook struct {
	Type string `json:"type"`
	Data struct {
		Account       string             `json:"account"`
		StatementItem TransactionRequest `json:"statementItem"`
	} `json:"data"`
}

func (m MonoWebhook) PrettyLog() string {
	tx := m.Data.StatementItem

	return fmt.Sprintf(
		`📩 Monobank Webhook Received
	Type:        %s
	Account:     %s
	TxID:        %s
	Amount:      %.2f
	Currency:    %d
	Balance:     %.2f
	MCC:         %d
	Description: %s
	Time:        %s`,
		m.Type,
		m.Data.Account,
		tx.ExternalId,
		float64(tx.Amount)/100,
		tx.CurrencyCode,
		float64(tx.Balance)/100,
		tx.MCC,
		tx.Description,
		time.Unix(tx.TransactionTime, 0).Format(time.RFC3339),
	)
}

type TransactionRequest struct {
	ExternalId      string `json:"id"`
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
