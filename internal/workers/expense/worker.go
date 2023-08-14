package expense

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
)

type Worker struct {
	repo     *store.DefaultRepo
	s3Client *s3.Client
}

func New(repo *store.DefaultRepo, s3Client *s3.Client) *Worker {
	return &Worker{
		repo:     repo,
		s3Client: s3Client,
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

	return w.s3Client.Upload(ctx, fmt.Sprintf("%s/%s", id, time.Now().Format("2006/01/02T15_04_05Z07_00")), bytes.NewReader(b))
}
