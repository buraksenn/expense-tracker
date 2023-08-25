package common

import "time"

type CommandType string

const (
	GetExpensesCommandType     CommandType = "GetExpensesCommand"
	RegisterExpenseCommandType CommandType = "RegisterExpenseCommand"
)

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

func GetCommandType(msg IncomingMessage) CommandType {
	if msg.Photo != "" {
		return RegisterExpenseCommandType
	}
	return GetExpensesCommandType
}
