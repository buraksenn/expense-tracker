package drive

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
)

type Client struct {
	svc *drive.Service
}

func NewClient(ctx context.Context) (*Client, error) {
	s, err := drive.NewService(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{svc: s}, nil
}

func (c *Client) Upload(content []byte) error {
	date := time.Now()
	fileID := fmt.Sprintf("%d-%d-%s", date.Month(), date.Day(), uuid.NewString())
	_, err := c.svc.Files.Create(&drive.File{
		Name: fileID,
	}).Media(bytes.NewReader(content)).Do()
	return err
}
