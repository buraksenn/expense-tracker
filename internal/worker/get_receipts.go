package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/buraksenn/expense-tracker/internal/common"
)

func (w *Worker) handleGetReceiptsCommand(ctx context.Context, msg *common.IncomingMessage) error {
	datePrefix := time.Now().Format("2006/01")
	if !msg.Date.IsZero() {
		datePrefix = msg.Date.Format("2006/01")
	}
	prefix := fmt.Sprintf("%s/%s", msg.User, datePrefix)

	images, err := w.s3Client.GetObjectsByPrefix(ctx, prefix)
	if err != nil {
		return fmt.Errorf("getting receipts from S3: %w", err)
	}

	if len(images) == 0 {
		w.outgoingMessageChan <- &common.OutgoingMessage{
			ChatID: msg.ChatID,
			Text:   "No receipts found for the specified year.",
		}
		return nil
	}

	w.outgoingMessageChan <- &common.OutgoingMessage{
		ChatID: msg.ChatID,
		Images: images,
	}

	return nil
}
