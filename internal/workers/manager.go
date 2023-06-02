package workers

type Manager struct {
}

// func GetCommandType(input string) (CommandType, error) {
// 	words := strings.Split(input, " ")
// 	switch len(words) {
// 	case ImageCommandLength:
// 		fallthrough
// 	case ExpenseCommandLength:
// 		log.Debug("Command length is valid for input: %s", input)
// 	default:
// 		return "", fmt.Errorf("invalid command")
// 	}

// 	switch strings.ToLower(input) {
// 	case "image", "i":
// 		return Image, nil
// 	case "expense", "e":
// 		return Expense, nil
// 	}
// 	return "", fmt.Errorf("invalid command")
// }
