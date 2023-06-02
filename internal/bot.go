package bot

import (
	"github.com/buraksenn/expense-tracker/pkg/drive"
	"github.com/buraksenn/expense-tracker/pkg/telegram"
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
}
