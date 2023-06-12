package expense

import (
	"strings"
	"time"
)

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

type GetExpensesCommand struct {
	StartDate time.Time
}

type RegisterExpenseCommand struct {
	ExpenseType   string
	Price         float32
	TaxPercentage int32
	Installment   int32
	Description   string
}
