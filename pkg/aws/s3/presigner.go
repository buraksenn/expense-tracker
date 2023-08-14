package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Presigner struct {
	svc *s3.PresignClient
}

func NewPresigner(c *Client) *Presigner {
	p := &Presigner{}
	p.svc = s3.NewPresignClient(c.svc)
	return p
}

// GetPresignedPutURL return a public http put url of s3 object with a period of expiresIn duration
func (c *Presigner) GetPresignedPutURL(ctx context.Context, bucket, key string, expiresIn time.Duration) (string, error) {
	req, err := c.svc.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiresIn
	})
	if err != nil {
		return "", fmt.Errorf("failed to get presigned put url, err: %v", err)
	}
	return req.URL, nil
}

// GetPresignedGetURL return a public http get url of s3 object with a period of expiresIn duration
func (c *Presigner) GetPresignedGetURL(ctx context.Context, bucket, key string, expiresIn time.Duration) (string, error) {
	req, err := c.svc.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = expiresIn
	})
	if err != nil {
		return "", fmt.Errorf("failed to get presigned get url, err: %v", err)
	}
	return req.URL, nil
}
