package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	DefaultBucket = "expensetracker-receipts"
)

type Client interface {
	Upload(ctx context.Context, key string, body io.Reader) error
}

type DefaultClient struct {
	svc *s3.Client
}

func NewClient(ctx context.Context) (*DefaultClient, error) {
	c := &DefaultClient{}
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	c.svc = s3.NewFromConfig(cfg)
	return c, nil
}

func (c *DefaultClient) Upload(ctx context.Context, key string, body io.Reader) error {
	_, err := c.svc.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(DefaultBucket),
		Key:         aws.String(key),
		ContentType: aws.String("image/jpeg"),
		Body:        body,
	})
	if err != nil {
		return fmt.Errorf("failed to upload s3 object with bucket: %s, key: %s, err: %v", DefaultBucket, key, err)
	}
	return nil
}
