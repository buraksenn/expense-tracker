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

	var s strings.Builder
	for _, e := range expenses {
		s.WriteString(fmt.Sprintf("Amount: %f, Tax: %f, CreatedAt: %d\n", e.Amount, e.Tax, e.CreatedAt))
	}
	w.outgoingMessageChan <- &common.OutgoingMessage{
		ChatID: msg.ChatID,
		Text:   s.String(),
	}
	logger.DebugC(ctx, "Expenses sent successfully.")
	return nil
}
