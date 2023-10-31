package worker

import (
	"context"
	"fmt"
	"strings"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

func (w *Worker) handleGetExpensesCommand(ctx context.Context, msg *common.IncomingMessage) error {
	expenses, err := w.repo.GetAllByID(ctx, msg.User)
	if err != nil {
		return fmt.Errorf("getting expenses: %w", err)
	}
	logger.Debug("Expenses: %+v", expenses)

	if len(expenses) == 0 {
		logger.DebugC(ctx, "No expenses found.")
		w.outgoingMessageChan <- &common.OutgoingMessage{
			ChatID: msg.ChatID,
			Text:   "No expenses found.",
		}
		return nil
	}

	var totalExpense float64
	var totalTax float64
	var s strings.Builder
	for _, e := range expenses {
		totalExpense += e.Amount
		totalTax += e.Tax
		s.WriteString(fmt.Sprintf("Amount: %f, Tax: %f, CreatedAt: %d\n", e.Amount, e.Tax, e.CreatedAt))
	}
	s.WriteString(fmt.Sprintf("Total Expense: %f, Total Tax: %f", totalExpense, totalTax))

	w.outgoingMessageChan <- &common.OutgoingMessage{
		ChatID: msg.ChatID,
		Text:   s.String(),
	}
	logger.DebugC(ctx, "Expenses sent successfully.")
	return nil
}
