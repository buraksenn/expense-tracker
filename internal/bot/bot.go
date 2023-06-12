package bot

import (
	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
	"github.com/buraksenn/expense-tracker/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	TelegramClient telegram.Client
	MessageChan    common.IncomingMessageChan
}

func New(t telegram.Client, messageChan common.IncomingMessageChan) *Bot {
	return &Bot{
		TelegramClient: t,
		MessageChan:    messageChan,
	}
}

func (b *Bot) Start() {
	ch, err := b.TelegramClient.GetUpdatesChan()
	if err != nil {
		logger.Fatal("Getting updates channel", err)
	}

	for update := range ch {
		if update.Message != nil {
			b.handleUpdate(&update)
		}
	}
}

func (b *Bot) handleUpdate(update *tgbotapi.Update) {
	msg := &common.IncomingMessage{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
		User:   update.Message.From.ID,
	}

	if len(update.Message.Photo) > 0 {
		msg.Photo = update.Message.Photo[0].FileID
	}

	b.MessageChan <- msg
}
