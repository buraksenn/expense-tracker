package worker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

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
	incomingMessages    common.IncomingMessageChan
	outgoingMessageChan common.OutgoingMessageChan
}

func New(repo store.Repo, s3Client s3.Client, textractClient textract.Client, i common.IncomingMessageChan, o common.OutgoingMessageChan) *Worker {
	return &Worker{
		repo:                repo,
		s3Client:            s3Client,
		textractClient:      textractClient,
		incomingMessages:    i,
		outgoingMessageChan: o,
	}
}

func (w *Worker) Start() {
	for msg := range w.incomingMessages {
		ctx := context.Background()
		logger.Debug("Received message: %+v", msg)

		switch common.GetCommandType(*msg) {
		case common.RegisterExpenseCommandType:
			w.handleRegisterExpenseCommand(ctx, msg)
		case common.GetExpensesCommandType:
			w.handleGetExpensesCommand(ctx, msg)
		}
	}
}

func (w *Worker) handleRegisterExpenseCommand(ctx context.Context, msg *common.IncomingMessage) error {
	s3Path, err := w.uploadPhoto(ctx, msg.User, msg.Photo)
	if err != nil {
		return fmt.Errorf("uploading photo: %w", err)
	}
	logger.Debug("Photo uploaded successfully. S3 path: %s", s3Path)

	tax, total, err := w.analyzePhoto(ctx, s3Path)
	if err != nil {
		return fmt.Errorf("analyzing photo: %w", err)
	}
	logger.Debug("Photo analyzed successfully. Tax and total: %f, %f", *tax, *total)
	w.outgoingMessageChan <- &common.OutgoingMessage{
		ChatID: msg.ChatID,
		Text:   fmt.Sprintf("Tax: %f, Total: %f", *tax, *total),
	}

	expense := &common.Expense{
		ID:        msg.User,
		Amount:    *total,
		Tax:       *tax,
		CreatedAt: time.Now().UTC().Unix(),
	}

	if err := w.registerExpense(ctx, expense); err != nil {
		return fmt.Errorf("registering expense: %w", err)
	}

	logger.Debug("Expense registered successfully. Expense: %+v", expense)
	return nil
}

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
	return nil
}

func (w *Worker) registerExpense(ctx context.Context, exp *common.Expense) error {
	err := w.repo.Put(ctx, exp)
	if err != nil {
		return fmt.Errorf("putting expense: %w", err)
	}
	return nil
}

func (w *Worker) uploadPhoto(ctx context.Context, id, link string) (string, error) {
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

func (w *Worker) analyzePhoto(ctx context.Context, s3Path string) (*float64, *float64, error) {
	exp, err := w.textractClient.QueryDocument(ctx, s3Path)
	if err != nil {
		return nil, nil, fmt.Errorf("analyzing document: %w", err)
	}

	tax, total, err := getTaxAndTotal(exp.Tax, exp.Total)
	if err != nil {
		return nil, nil, fmt.Errorf("getting tax and total: %w", err)
	}

	return &tax, &total, err
}

func getTaxAndTotal(t, tot string) (float64, float64, error) {
	t = omitEverythingExceptDigits(t)
	tot = omitEverythingExceptDigits(tot)

	tax, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing tax: %w", err)
	}
	total, err := strconv.ParseFloat(tot, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing total: %w", err)
	}
	// Tax and total strings do not have decimal point and comma in them so we need to divide them by 100
	return tax / 100, total / 100, nil
}

func omitEverythingExceptDigits(s string) string {
	var b bytes.Buffer
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
