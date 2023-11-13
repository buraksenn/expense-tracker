package textract

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
	"github.com/buraksenn/expense-tracker/pkg/logger"
)

const (
	TAX_QUERY   = "What is the value in the same line of TOPKDV?"
	TOTAL_QUERY = "What is the value in the same line of TOPLAM?"
)

type AnalyzedExpense struct {
	Tax   string `json:"tax"`
	Total string `json:"total"`
}

type Client interface {
	QueryDocument(ctx context.Context, path string) (*AnalyzedExpense, error)
	AnalyzeExpense(ctx context.Context, path string) (*AnalyzedExpense, error)
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

func (c *DefaultClient) AnalyzeExpense(ctx context.Context, path string) (*AnalyzedExpense, error) {
	inp := &textract.AnalyzeExpenseInput{
		Document: &types.Document{
			S3Object: &types.S3Object{
				Bucket: aws.String(s3.DefaultBucket),
				Name:   aws.String(path),
			},
		},
	}

	out, err := c.cl.AnalyzeExpense(ctx, inp)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, fmt.Errorf("output is nil")
	}
	tax := ""
	total := ""
	subTotal := ""
	for _, s := range out.ExpenseDocuments[0].SummaryFields {
		if *s.Type.Text == "TAX" {
			tax = *s.ValueDetection.Text
		}
		if *s.Type.Text == "TOTAL" {
			total = *s.ValueDetection.Text
		}
		if *s.Type.Text == "SUBTOTAL" {
			subTotal = *s.ValueDetection.Text
		}
	}

	logger.DebugC(ctx, "AnalyzeExpense Tax: %s, Total: %s, SubTotal: %s", tax, total, subTotal)

	if tax == "" {
		tax = subTotal
	}

	return &AnalyzedExpense{
		Tax:   tax,
		Total: total,
	}, nil
}

// Not used
func (c *DefaultClient) QueryDocument(ctx context.Context, path string) (*AnalyzedExpense, error) {
	inp := &textract.AnalyzeDocumentInput{
		Document: &types.Document{
			S3Object: &types.S3Object{
				Bucket: aws.String(s3.DefaultBucket),
				Name:   aws.String(path),
			},
		},
		FeatureTypes: []types.FeatureType{
			types.FeatureTypeQueries,
		},
		QueriesConfig: &types.QueriesConfig{
			Queries: []types.Query{
				{
					Text:  aws.String(TAX_QUERY),
					Alias: aws.String("Q_TOPKDV"),
				},
				{
					Text:  aws.String(TOTAL_QUERY),
					Alias: aws.String("Q_TOPLAM"),
				},
			},
		},
	}

	out, err := c.cl.AnalyzeDocument(ctx, inp)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, fmt.Errorf("output is nil")
	}

	queryAndAnswerBlocks := map[string]types.Block{}
	for _, block := range out.Blocks {
		if block.BlockType == types.BlockTypeQuery || block.BlockType == types.BlockTypeQueryResult {
			queryAndAnswerBlocks[*block.Id] = block
		}
	}

	queryAndAnswers := map[string]string{}
	for _, q := range queryAndAnswerBlocks {
		if q.BlockType != types.BlockTypeQuery {
			continue
		}
		if len(q.Relationships) != 1 {
			return nil, fmt.Errorf("query block has more than one relationship")
		}
		rel := q.Relationships[0]
		if rel.Type == types.RelationshipTypeAnswer {
			if len(rel.Ids) != 1 {
				return nil, fmt.Errorf("query block has more than one answer")
			}
			id := rel.Ids[0]
			answer, ok := queryAndAnswerBlocks[id]
			if !ok {
				return nil, fmt.Errorf("answer block not found")
			}
			queryAndAnswers[*q.Query.Text] = *answer.Text
		}
	}

	return &AnalyzedExpense{
		Tax:   queryAndAnswers[TAX_QUERY],
		Total: queryAndAnswers[TOTAL_QUERY],
	}, nil
}
