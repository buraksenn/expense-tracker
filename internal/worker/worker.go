package worker

import (
	"context"
	"fmt"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/store"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
	"github.com/buraksenn/expense-tracker/pkg/aws/textract"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

type Worker struct {
	repo                store.Repo
	s3Client            s3.Client
	textractClient      textract.Client
	outgoingMessageChan common.OutgoingMessageChan
}

func New(repo store.Repo, s3Client s3.Client, textractClient textract.Client, o common.OutgoingMessageChan) *Worker {
	return &Worker{
		repo:                repo,
		s3Client:            s3Client,
		textractClient:      textractClient,
		outgoingMessageChan: o,
	}
}

func (w *Worker) HandleCommand(ctx context.Context, msg *common.IncomingMessage) error {
	logger.DebugC(ctx, "Received message: %+v", msg)

	switch common.GetCommandType(*msg) {
	case common.RegisterExpenseCommandType:
		return w.handleRegisterExpenseCommand(ctx, msg)
	case common.GetExpensesCommandType:
		return w.handleGetExpensesCommand(ctx, msg)
	default:
		return fmt.Errorf("unrecognized command type: %s", msg.Text)
	}
}
