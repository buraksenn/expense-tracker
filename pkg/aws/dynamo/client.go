package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Client interface {
	GetItem(ctx context.Context, inp *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	GetItems(ctx context.Context, inp *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	PutItem(ctx context.Context, inp *dynamodb.PutItemInput) error
}

type DefaultClient struct {
	DB *dynamodb.Client
}

func NewClient(ctx context.Context) (*DefaultClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &DefaultClient{
		DB: dynamodb.NewFromConfig(cfg),
	}, nil
}

func (c *DefaultClient) GetItem(ctx context.Context, inp *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	out, err := c.DB.GetItem(ctx, inp)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *DefaultClient) GetItems(ctx context.Context, inp *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	out, err := c.DB.Query(ctx, inp)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *DefaultClient) PutItem(ctx context.Context, inp *dynamodb.PutItemInput) error {
	_, err := c.DB.PutItem(ctx, inp)
	if err != nil {
		return err
	}
	return nil
}
