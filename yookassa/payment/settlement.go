// Package yoopayment describes all the necessary entities for working with YooMoney Payments.
package yoopayment

// Information about money distribution.
type Settlement struct {
	// Transaction type. Fixed value: payout — payout to seller.
	Type string `json:"type,omitempty"`

	// Amount of seller’s remuneration.
	Amount Amount `json:"amount,omitempty"`
}
