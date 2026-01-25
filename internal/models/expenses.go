package models

import "time"

const (
	TypeIncome  = "Income"
	TypeExpense = "Expense"
)

type Expense struct {
	Id              int64     `json:"id"`
	Currency        string    `json:"currency"`
	Amount          int64     `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	Descriptions    string    `json:"descriptions"`
	DateOfExpense   time.Time `json:"date_of_expense"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	UserId          int64     `json:"user_id"`
	CategoryId      int64     `json:"category_id"`
}
