package workers

import (
	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/internal/store"
	"github.com/buraksenn/expense-tracker/internal/workers/expense"
	"github.com/buraksenn/expense-tracker/pkg/aws/s3"
)

type NewManagerInput struct {
	MessageChan common.IncomingMessageChan
	S3Client    *s3.Client
	Repo        *store.DefaultRepo
}

type Manager struct {
	MessageChan   common.IncomingMessageChan
	ExpenseWorker *expense.Worker
}

func New(inp *NewManagerInput) *Manager {
	return &Manager{
		MessageChan:   inp.MessageChan,
		ExpenseWorker: expense.New(inp.Repo, inp.S3Client),
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
