package s3

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	DefaultBucket = "expensetracker-receipts"
)

type Client struct {
	svc           *s3.Client
	presignClient *s3.PresignClient
}

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	c := &Client{
		svc: s3.NewFromConfig(cfg),
	}
	c.presignClient = s3.NewPresignClient(c.svc)

	return c, nil
}

func (c *Client) Upload(ctx context.Context, key string, body io.Reader) error {
	_, err := c.svc.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(DefaultBucket),
		Key:         aws.String(key),
		ContentType: aws.String("image/jpeg"),
		Body:        body,
	})
	if err != nil {
		return fmt.Errorf("uploading s3 object with bucket: %s, key: %s, err: %v", DefaultBucket, key, err)
	}
	return nil
}

func (c *Client) GetObjectsByPrefix(ctx context.Context, prefix string) ([]string, error) {
	// First, get all object keys
	var objectKeys []string
	paginator := s3.NewListObjectsV2Paginator(c.svc, &s3.ListObjectsV2Input{
		Bucket: aws.String(DefaultBucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("listing objects with prefix %s: %v", prefix, err)
		}

		for _, object := range page.Contents {
			objectKeys = append(objectKeys, *object.Key)
		}
	}

	// Get presigned urls for each object concurrently
	urls := make([]string, len(objectKeys))
	errChan := make(chan error, len(objectKeys))
	var wg sync.WaitGroup

	for i, key := range objectKeys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			req, err := c.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
				Bucket: aws.String(DefaultBucket),
				Key:    aws.String(key),
			}, func(opts *s3.PresignOptions) {
				opts.Expires = time.Hour * 24
			})
			if err != nil {
				errChan <- fmt.Errorf("presigning URL for object %s: %v", key, err)
				return
			}
			urls[i] = req.URL
		}(i, key)
	}

	wg.Wait()
	close(errChan)

	if err := <-errChan; err != nil {
		return nil, err
	}

	return urls, nil
}
