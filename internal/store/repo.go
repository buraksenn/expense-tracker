package store

import (
	"context"
	"fmt"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/aws/dynamo"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	TableName = aws.String("expenses")
)

type Repo interface {
	Put(ctx context.Context, expense *common.Expense) error
	GetAllByID(ctx context.Context, chatID string) ([]common.Expense, error)
}

type DefaultRepo struct {
	c dynamo.Client
}

func NewRepo(cl dynamo.Client) *DefaultRepo {
	return &DefaultRepo{
		c: cl,
	}
}

func (r *DefaultRepo) Put(ctx context.Context, expense *common.Expense) error {
	av, err := attributevalue.MarshalMap(expense)
	if err != nil {
		return fmt.Errorf("failed to marshal expense, err: %v", err)
	}

	return r.c.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: TableName,
		Item:      av,
	})
}

func (r *DefaultRepo) GetAllByID(ctx context.Context, chatID string) ([]common.Expense, error) {
	keyEx := expression.Key("id").Equal(expression.Value(chatID))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build expression, err: %v", err)
	}

	out, err := r.c.GetItems(ctx, &dynamodb.QueryInput{
		TableName:                 TableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to query dynamodb, err: %v", err)
	}

	var e []common.Expense
	if err := attributevalue.UnmarshalListOfMaps(out.Items, &e); err != nil {
		return nil, fmt.Errorf("failed to unmarshall expenses, err: %v", err)
	}

	return e, nil
}
