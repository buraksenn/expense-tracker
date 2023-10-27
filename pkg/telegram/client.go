package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DefaultClient struct {
	bot *tgbotapi.BotAPI
}

func New(token string) (*DefaultClient, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &DefaultClient{bot: bot}, nil
}

func (cl *DefaultClient) SendMessage(chatID int64, text string) error {
	return cl.sendMessageInternal(chatID, text, 0)
}

func (cl *DefaultClient) sendMessageInternal(chatID int64, text string, repliedMessageID int) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = repliedMessageID
	_, err := cl.bot.Send(msg)
	return err
}

func (cl *DefaultClient) GetFileLink(id string) (string, error) {
	fileConfig := tgbotapi.FileConfig{
		FileID: id,
	}

	file, err := cl.bot.GetFile(fileConfig)
	if err != nil {
		return "", fmt.Errorf("getting file: %w", err)
	}

	return file.Link(cl.bot.Token), nil
}

func (cl *DefaultClient) SendImage(chatID int64, url string) error {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(url))
	_, err := cl.bot.Send(photo)
	return err
}

func (cl *DefaultClient) SendImageBatch(chatID int64, urls []string) error {
	// TODO: use goroutine to send images concurrently or use telegram API to send multiple images at once
	for _, url := range urls {
		if err := cl.SendImage(chatID, url); err != nil {
			return err
		}
	}
	return nil
}
