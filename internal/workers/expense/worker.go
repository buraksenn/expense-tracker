package expense

import (
	"context"
	"database/sql"
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

func (w *Worker) GetExpenses(ctx context.Context, cmd GetExpensesCommand) ([]*repository.Expense, error) {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return nil, fmt.Errorf("invalid userID")
	}

	return w.Repository.GetExpenses(ctx, &repository.GetExpensesParams{
		UserID:      userID,
		CreatedAt:   cmd.StartDate,
		CreatedAt_2: time.Now(),
	})
}

func (w *Worker) GetExpensesSummary(ctx context.Context, cmd GetExpensesCommand) ([]*repository.GetExpensesSummaryRow, error) {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return nil, fmt.Errorf("invalid userID")
	}

	return w.Repository.GetExpensesSummary(ctx, &repository.GetExpensesSummaryParams{
		UserID:      userID,
		CreatedAt:   cmd.StartDate,
		CreatedAt_2: time.Now(),
	})
}

func (w *Worker) CreateExpense(ctx context.Context, cmd RegisterExpenseCommand) error {
	userID, ok := ctx.Value("userID").(int32)
	if !ok || userID == 0 {
		return fmt.Errorf("invalid userID")
	}

	installment := sql.NullInt32{
		Int32: int32(cmd.Installment),
		Valid: cmd.Installment > 0,
	}

	installmentEndDate := sql.NullTime{
		Time:  time.Now().Add(time.Duration(cmd.Installment) * 30 * 24 * time.Hour),
		Valid: cmd.Installment > 0,
	}

	_, err := w.Repository.CreateExpense(ctx, &repository.CreateExpenseParams{
		UserID:             userID,
		Description:        cmd.Description,
		Type:               string(GetType(cmd.ExpenseType)),
		Price:              cmd.Price,
		TaxPercentage:      cmd.TaxPercentage,
		Installment:        installment,
		InstallmentEndDate: installmentEndDate,
	})
	return err
}
