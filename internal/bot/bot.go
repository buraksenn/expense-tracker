package bot

import (
	"fmt"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type telegramClient interface {
	SendMessage(chatID int64, text string) error
	SendImage(chatID int64, url string) error
	GetFileLink(fileID string) (string, error)
}

type Bot struct {
	telegramCl telegramClient
	inChan     common.OutgoingMessageChan
	DoneChan   chan struct{}
}

func New(t telegramClient, i common.OutgoingMessageChan, doneC chan struct{}) *Bot {
	return &Bot{
		telegramCl: t,
		inChan:     i,
		DoneChan:   doneC,
	}
}

func (b *Bot) Start() {
	go b.handleOutgoing()
}

func (b *Bot) Stop() {
	logger.Debug("Stopping telegram bot...")
	close(b.inChan)
}

func (b *Bot) PrepareMessage(update *tgbotapi.Update) (*common.IncomingMessage, error) {
	if update == nil || update.Message == nil {
		return nil, fmt.Errorf("update or message is nil")
	}

	msg := &common.IncomingMessage{
		ChatID: update.Message.Chat.ID,
		User:   fmt.Sprint(update.Message.From.ID),
		Text:   update.Message.Text,
	}

	if len(update.Message.Photo) > 0 {
		fileID := update.Message.Photo[len(update.Message.Photo)-1].FileID
		link, err := b.telegramCl.GetFileLink(fileID)
		if err != nil {
			return nil, fmt.Errorf("getting file link: %w", err)
		}
		msg.Photo = link
	}

	return msg, nil
}

func (b *Bot) handleOutgoing() {
	for msg := range b.inChan {
		logger.Debug("Sending message to chat: %d, text: %s", msg.ChatID, msg.Text)
		if err := b.telegramCl.SendMessage(msg.ChatID, msg.Text); err != nil {
			logger.Error("Sending message to chat: %d, err: %v", msg.ChatID, err)
		}
	}
	logger.Debug("Outgoing message channel closed")
	b.DoneChan <- struct{}{}
}
