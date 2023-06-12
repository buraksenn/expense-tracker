package workers

import "github.com/buraksenn/expense-tracker/internal/common"

type Manager struct {
	MessageChan common.IncomingMessageChan
}

func New(MessageChan common.IncomingMessageChan) *Manager {
	return &Manager{
		MessageChan: MessageChan,
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
