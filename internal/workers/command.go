package workers

import "strings"

type CommandType string

const (
	GetExpenses     CommandType = "GetExpenses"
	GetImages       CommandType = "GetImages"
	RegisterExpense CommandType = "RegisterExpense"
	RegisterImage   CommandType = "RegisterImage"
)

func GetCommandType(s string) CommandType {
	switch strings.ToLower(s) {
	case "ge":
		return GetExpenses
	case "gi":
		return GetImages
	case "e":
		return RegisterExpense
	default:
		return ""
	}
}
