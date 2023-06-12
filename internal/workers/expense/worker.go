package expense

import (
	"context"
	"fmt"
	"time"

	"github.com/buraksenn/expense-tracker/internal/db/repository"
	"github.com/buraksenn/expense-tracker/pkg/spreadsheet"
)

type Worker struct {
	SpreadsheetClient spreadsheet.Client
	Repository        repository.Querier
}

func New(s spreadsheet.Client, r repository.Querier) *Worker {
	return &Worker{
		SpreadsheetClient: s,
		Repository:        r,
	}
}

func (w *Worker) Run() {

}

func (w *Worker) GetExpenses(ctx context.Context, command GetExpensesCommand) ([]*repository.Expense, error) {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return nil, fmt.Errorf("invalid userID")
	}

	return w.Repository.GetExpenses(ctx, &repository.GetExpensesParams{
		UserID:      userID,
		CreatedAt:   command.StartDate,
		CreatedAt_2: time.Now(),
	})
}

func (w *Worker) GetExpensesSummary(ctx context.Context, command GetExpensesCommand) ([]*repository.GetExpensesSummaryRow, error) {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return nil, fmt.Errorf("invalid userID")
	}

	return w.Repository.GetExpensesSummary(ctx, &repository.GetExpensesSummaryParams{
		UserID:      userID,
		CreatedAt:   command.StartDate,
		CreatedAt_2: time.Now(),
	})
}

func (w *Worker) CreateExpense(ctx context.Context, cmd RegisterExpenseCommand) error {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return fmt.Errorf("invalid userID")
	}

	_, err := w.Repository.CreateExpense(ctx, &repository.CreateExpenseParams{
		UserID:        userID,
		Description:   cmd.Description,
		Type:          string(GetType(cmd.ExpenseType)),
		Price:         cmd.Price,
		TaxPercentage: cmd.TaxPercentage,
	})
	return err
}
