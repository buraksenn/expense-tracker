package bot

import (
	"github.com/buraksenn/expense-tracker/pkg/drive"
	"github.com/buraksenn/expense-tracker/pkg/logger"
	"github.com/buraksenn/expense-tracker/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
}

type Bot struct {
	TelegramClient telegram.Client
	Drive          drive.Client
}

func New(t telegram.Client, d drive.Client) *Bot {
	return &Bot{
		TelegramClient: t,
		Drive:          d,
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

}
