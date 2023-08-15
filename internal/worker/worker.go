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
	"github.com/buraksenn/expense-tracker/pkg/aws/textract"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

type Worker struct {
	repo             store.Repo
	s3Client         s3.Client
	textractClient   textract.Client
	incomingMessages common.IncomingMessageChan
}

func New(repo store.Repo, s3Client s3.Client, textractClient textract.Client, c common.IncomingMessageChan) *Worker {
	return &Worker{
		repo:             repo,
		s3Client:         s3Client,
		textractClient:   textractClient,
		incomingMessages: c,
	}
}

func (w *Worker) Start() {
	ctx := context.Background()
	for msg := range w.incomingMessages {
		logger.Info("Received message: %+v", msg)
		switch GetCommandType(*msg) {
		case common.RegisterExpenseCommandType:
			s3Path, err := w.UploadPhoto(ctx, msg.User, msg.Photo)
			if err != nil {
				logger.Error("Uploading photo: %v", err)
				continue
			}
			logger.Info("Photo uploaded successfully.")
			if err := w.AnalyzePhoto(ctx, s3Path); err != nil {
				logger.Error("Analyzing photo: %v", err)
				continue
			}
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
	return nil
}

func (w *Worker) UploadPhoto(ctx context.Context, id, link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("parsing url: %w", err)
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Do(&http.Request{
		Method: http.MethodGet,
		URL:    u,
	})
	if err != nil {
		return "", fmt.Errorf("downloading file: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("downloading file: status code %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	s3Path := fmt.Sprintf("%s/%s.jpeg", id, time.Now().Format("2006/01/02T15_04_05Z07_00"))
	return s3Path, w.s3Client.Upload(ctx, s3Path, bytes.NewReader(b))
}

func (w *Worker) AnalyzePhoto(ctx context.Context, s3Path string) error {
	out, err := w.textractClient.AnalyzeDocument(ctx, s3Path)
	if err != nil {
		return fmt.Errorf("analyzing document: %w", err)
	}
	logger.Info("Analyzed document: %+v", out)
	for _, block := range out.Blocks {
		if block.Text != nil {
			logger.Info("Found text: %s", *block.Text)
		}
	}

	return nil
}

func GetCommandType(msg common.IncomingMessage) common.CommandType {
	if msg.Photo != "" {
		return common.RegisterExpenseCommandType
	}
	return common.GetExpensesCommandType
}
