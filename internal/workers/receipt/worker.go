package receipt

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/buraksenn/expense-tracker/pkg/drive"
)

type Worker struct {
	drive  drive.Client
	httpCl *http.Client
}

func New(d drive.Client) *Worker {
	return &Worker{
		drive: d,
		// TODO cleanhttp client
		httpCl: http.DefaultClient,
	}
}

func (w *Worker) UploadReceipt(ctx context.Context, cmd UploadReceiptCommand) error {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return fmt.Errorf("invalid userID")
	}

	u, err := url.Parse(cmd.Link)
	if err != nil {
		return fmt.Errorf("parsing url: %w", err)
	}

	resp, err := w.httpCl.Do(&http.Request{
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

	return w.drive.Upload(string(userID), b)
}

func (w *Worker) GetReceipts(ctx context.Context, cmd GetReceiptsCommand) ([]byte, error) {
	return nil, nil
}
