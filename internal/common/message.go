package common

type IncomingMessage struct {
	ChatID int64
	User   string
	Text   string
	Photo  string
}

type OutgoingMessage struct {
	ChatID int64
	Text   string
}
type OutgoingMessageChan chan *OutgoingMessage
