package telegram

import (
	"errors"
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func New(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Client{bot: bot}, nil
}

func (cl *Client) SendMessage(chatID int64, text string) error {
	return cl.sendMessageInternal(chatID, text, 0)
}

func (cl *Client) sendMessageInternal(chatID int64, text string, repliedMessageID int) error {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyToMessageID = repliedMessageID
	_, err := cl.bot.Send(msg)
	return err
}

func (cl *Client) GetFileLink(id string) (string, error) {
	fileConfig := tgbotapi.FileConfig{
		FileID: id,
	}

	file, err := cl.bot.GetFile(fileConfig)
	if err != nil {
		return "", fmt.Errorf("getting file: %w", err)
	}

	return file.Link(cl.bot.Token), nil
}

func (cl *Client) SendImage(chatID int64, url string) error {
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(url))
	_, err := cl.bot.Send(photo)
	return err
}

func (cl *Client) BatchSendImage(chatID int64, urls []string) error {
	var wg sync.WaitGroup
	var err error
	var errMu sync.Mutex
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if e := cl.SendImage(chatID, url); err != nil {
				errMu.Lock()
				err = errors.Join(err, e)
				errMu.Unlock()
			}
		}(url)
	}
	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
