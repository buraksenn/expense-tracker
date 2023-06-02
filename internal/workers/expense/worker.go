package expense

import (
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
