package common

type CommandType string

const (
	GetExpensesCommandType     CommandType = "GetExpensesCommand"
	GetReceiptsCommandType     CommandType = "GetReceiptsCommand"
	RegisterExpenseCommandType CommandType = "RegisterExpenseCommand"
)

func GetCommandType(msg IncomingMessage) CommandType {
	if msg.Photo != "" {
		return RegisterExpenseCommandType
	}
	if msg.Text == "receipts" {
		return GetReceiptsCommandType
	}

	return GetExpensesCommandType
}
