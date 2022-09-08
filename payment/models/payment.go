package models

type Payment struct {
	ID       int     `json:"id"`
	Ref      string  `json:"reference"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
