package textract

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
)

type Client interface {
	AnalyzeDocument(ctx context.Context, path string) (*textract.AnalyzeDocumentOutput, error)
}

type DefaultClient struct {
	cl *textract.Client
}

func NewDefaultClient(ctx context.Context) (*DefaultClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &DefaultClient{
		cl: textract.NewFromConfig(cfg),
	}, nil
}

func (c *DefaultClient) AnalyzeDocument(ctx context.Context, path string) (*textract.AnalyzeDocumentOutput, error) {
	inp := &textract.AnalyzeDocumentInput{
		Document: &types.Document{
			S3Object: &types.S3Object{
				Bucket: aws.String(s3.DefaultBucket),
				Name:   aws.String(path),
			},
		},
		FeatureTypes: []types.FeatureType{types.FeatureTypeQueries},
		QueriesConfig: &types.QueriesConfig{
			Queries: []types.Query{
				{
					Text: aws.String("What is TOPKDV"),
				},
				{
					Text: aws.String("What is TOPLAM"),
				},
			},
		},
	}

	out, err := c.cl.AnalyzeDocument(ctx, inp)
	if err != nil {
		return nil, err
	}

	return out, nil
}
