package worker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/store"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

type Worker struct {
	repo             *store.DefaultRepo
	s3Client         *s3.Client
	incomingMessages common.IncomingMessageChan
}

func New(repo *store.DefaultRepo, s3Client *s3.Client, c common.IncomingMessageChan) *Worker {
	return &Worker{
		repo:             repo,
		s3Client:         s3Client,
		incomingMessages: c,
	}
}

func (w *Worker) Start() {
	for msg := range w.incomingMessages {
		logger.Info("Received message: %+v", msg)
		switch GetCommandType(*msg) {
		case common.RegisterExpenseCommandType:
			err := w.UploadPhoto(context.Background(), msg.User, msg.Photo)
			if err != nil {
				logger.Error("Uploading photo: %v", err)
				continue
			}
			logger.Info("Photo uploaded successfully.")
		case common.GetExpensesCommandType:
			logger.Info("GetExpensesCommandType not implemented yet.")
		}
	}
}

func (w *Worker) GetExpenses(ctx context.Context, cmd common.GetExpensesCommand) ([]*common.Expense, error) {
	expenses, err := w.repo.GetAllByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}
	// TODO Format expenses
	return expenses, nil
}

func (w *Worker) RegisterExpense(ctx context.Context, cmd common.RegisterExpenseCommand) error {
	err := w.repo.Put(ctx, &common.Expense{
		ID:        cmd.ID,
		Amount:    cmd.Amount,
		Tax:       cmd.Tax,
		CreatedAt: time.Now().UTC().Unix(),
	})
	if err != nil {
		return fmt.Errorf("putting expense: %w", err)
	}
	return w.UploadPhoto(ctx, cmd.ID, cmd.Photo)
}

func (w *Worker) UploadPhoto(ctx context.Context, id, link string) error {
	u, err := url.Parse(link)
	if err != nil {
		return fmt.Errorf("parsing url: %w", err)
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
	})
	if err != nil {
		return fmt.Errorf("downloading file: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("downloading file: status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	return w.s3Client.Upload(ctx, fmt.Sprintf("%s/%s.jpeg", id, time.Now().Format("2006/01/02T15_04_05Z07_00")), bytes.NewReader(b))
}

func GetCommandType(msg common.IncomingMessage) common.CommandType {
	if msg.Photo != "" {
		return common.RegisterExpenseCommandType
	}
	return common.GetExpensesCommandType
}
