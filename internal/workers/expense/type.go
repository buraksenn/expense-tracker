package expense

import "strings"

type ExpenseType string

const (
	Food       ExpenseType = "Food"
	Groceries  ExpenseType = "Groceries"
	Fuel       ExpenseType = "Fuel"
	Tech       ExpenseType = "Tech"
	House      ExpenseType = "House"
	HealthCare ExpenseType = "HealthCare"
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

type Command struct {
	ExpenseType   ExpenseType
	Price         float32
	TaxPercentage int
	Installment   int
	Description   string
}
