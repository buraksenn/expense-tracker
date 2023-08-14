package common

type IncomingMessage struct {
	ChatID int64
	User   int64
	Text   string
	Photo  string
}

type IncomingMessageChan chan *IncomingMessage
