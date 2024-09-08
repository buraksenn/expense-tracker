package worker

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

func (w *Worker) handleGetExpensesCommand(ctx context.Context, msg *common.IncomingMessage) error {
	var expenses []common.Expense
	var err error

	if msg.Text == "all" {
		expenses, err = w.repo.GetAllByID(ctx, msg.User)
	} else {
		expenses, err = w.repo.GetMonthExpenses(ctx, msg.User, msg.Date)
	}
	if err != nil {
		return fmt.Errorf("getting expenses: %w", err)
	}

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

	for i, e := range expenses {
		totalExpense += e.Amount
		totalTax += e.Tax

		date := time.Unix(e.CreatedAt, 0).Format(time.RFC3339)
		s.WriteString(fmt.Sprintf("Amount: %.2f, Tax: %.2f, CreatedAt: %s\n", e.Amount, e.Tax, date))

		// Send the message every 10 expenses
		if i%10 == 0 {
			w.outgoingMessageChan <- &common.OutgoingMessage{
				ChatID: msg.ChatID,
				Text:   s.String(),
			}
			s.Reset()
		}
	}

	w.outgoingMessageChan <- &common.OutgoingMessage{
		ChatID: msg.ChatID,
		Text:   fmt.Sprintf("Total Expense: %f, Total Tax: %f", totalExpense, totalTax),
	}

	logger.DebugC(ctx, "Expenses sent successfully.")
	return nil
}
