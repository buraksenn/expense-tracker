// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package repository

import (
	"context"
)

type Querier interface {
	CreateExpense(ctx context.Context, arg *CreateExpenseParams) (*Expense, error)
	GetExpenses(ctx context.Context, arg *GetExpensesParams) ([]*Expense, error)
	GetExpensesSummary(ctx context.Context, arg *GetExpensesSummaryParams) ([]*GetExpensesSummaryRow, error)
	GetUser(ctx context.Context, id int32) (*User, error)
}

var _ Querier = (*Queries)(nil)
