package common

import "time"

type GetExpensesCommand struct {
	ID        string
	StartDate time.Time
}

type RegisterExpenseCommand struct {
	ID     string
	Amount int
	Tax    int
	Photo  string
}
