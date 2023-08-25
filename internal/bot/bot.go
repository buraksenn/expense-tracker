package bot

import (
	"fmt"

	"github.com/buraksenn/expense-tracker/internal/common"
	"github.com/buraksenn/expense-tracker/pkg/logger"
	"github.com/buraksenn/expense-tracker/pkg/telegram"
)

const (
	concurrency = 10
)

type Bot struct {
	telegramClient      telegram.Client
	incomingMessageChan common.IncomingMessageChan
	outgoingMessageChan common.OutgoingMessageChan
}

func New(t telegram.Client, i common.IncomingMessageChan, o common.OutgoingMessageChan) *Bot {
	return &Bot{
		telegramClient:      t,
		incomingMessageChan: i,
		outgoingMessageChan: o,
	}
}

func (b *Bot) Start() {
	go b.handleIncoming()
	go b.handleOutgoing()

}

func (b *Bot) handleIncoming() {
	ch, err := b.telegramClient.GetUpdatesChan()
	if err != nil {
		logger.Fatal("Getting updates channel", err)
	}

	for update := range ch {
		if update.Message != nil {
			msg := &common.IncomingMessage{
				ChatID: update.Message.Chat.ID,
				User:   fmt.Sprint(update.Message.From.ID),
				Text:   update.Message.Text,
			}

			if len(update.Message.Photo) > 0 {
				fileID := update.Message.Photo[len(update.Message.Photo)-1].FileID
				link, err := b.telegramClient.GetFileLink(fileID)
				if err != nil {
					// TODO handle error properly
					logger.Error("Getting file link for msg: %+v, err: %+v", msg, err)
				}
				msg.Photo = link
			}

			b.incomingMessageChan <- msg
		}
	}
}

func (b *Bot) handleOutgoing() {
	for msg := range b.outgoingMessageChan {
		if err := b.telegramClient.SendMessage(msg.ChatID, msg.Text); err != nil {
			// TODO handle error properly
			logger.Error("Sending message to chat: %d, err: %v", msg.ChatID, err)
		}
	}

}
