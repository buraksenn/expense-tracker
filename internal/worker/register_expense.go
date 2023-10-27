package worker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"unicode"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

func (w *Worker) handleRegisterExpenseCommand(ctx context.Context, msg *common.IncomingMessage) error {
	s3Path, err := w.uploadPhoto(ctx, msg.User, msg.Photo)
	if err != nil {
		return fmt.Errorf("uploading photo: %w", err)
	}
	logger.DebugC(ctx, "Photo uploaded successfully. S3 path: %s", s3Path)

	tax, total, err := w.analyzePhoto(ctx, s3Path)
	if err != nil {
		return fmt.Errorf("analyzing photo: %w", err)
	}
	logger.DebugC(ctx, "Photo analyzed successfully. Tax and total: %f, %f", *tax, *total)
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

	logger.DebugC(ctx, "Expense registered successfully. Expense: %+v", expense)
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
