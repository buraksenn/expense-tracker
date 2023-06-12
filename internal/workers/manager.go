package workers

import (
	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/db/repository"
	"github.com/buraksenn/expense-tracker/internal/workers/expense"
	"github.com/buraksenn/expense-tracker/internal/workers/receipt"
	"github.com/buraksenn/expense-tracker/pkg/drive"
	"github.com/buraksenn/expense-tracker/pkg/spreadsheet"
)

type NewManagerInput struct {
	MessageChan       common.IncomingMessageChan
	DriveClient       drive.Client
	SpreadsheetClient spreadsheet.Client
	Repository        repository.Querier
}

type Manager struct {
	MessageChan   common.IncomingMessageChan
	ReceiptWorker *receipt.Worker
	ExpenseWorker *expense.Worker
}

func New(inp *NewManagerInput) *Manager {
	return &Manager{
		MessageChan:   inp.MessageChan,
		ReceiptWorker: receipt.New(inp.DriveClient),
		ExpenseWorker: expense.New(inp.SpreadsheetClient, inp.Repository),
	}
}

func (m *Manager) Run() {

	message := <-m.MessageChan
	if message.Photo != "" {
		// TODO: register receipt, pass it to worker
	}

	command := GetCommandType(message.Text)
	if command == "" {
		// TODO: send error message
	}

}
