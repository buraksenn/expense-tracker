package common

import "strings"

type Expense struct {
	ID        string      `json:"id"`
	Type      ExpenseType `json:"type"`
	Amount    int         `json:"amount"`
	CreatedAt int64       `json:"created_at"`
	Tax       int         `json:"tax"`
}

type ExpenseType string

const (
	Food       ExpenseType = "Food"
	Groceries  ExpenseType = "Groceries"
	Fuel       ExpenseType = "Fuel"
	Tech       ExpenseType = "Tech"
	House      ExpenseType = "House"
	HealthCare ExpenseType = "Healthcare"
	Other      ExpenseType = "Other"
)

func GetType(s string) ExpenseType {
	switch strings.ToLower(s) {
	case "food", "f":
		return Food
	case "groceries", "g":
		return Groceries
	case "fuel", "fl":
		return Fuel
	case "tech", "t":
		return Tech
	case "house", "h":
		return House
	case "health", "hc":
		return HealthCare
	default:
		return Other
	}
}
