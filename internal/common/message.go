package common

import "time"

type IncomingMessage struct {
	ChatID int64
	Date   time.Time
	User   string
	Text   string
	Photo  string
}

type OutgoingMessage struct {
	ChatID int64
	Text   string
}
type OutgoingMessageChan chan *OutgoingMessage
