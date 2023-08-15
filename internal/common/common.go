package common

type IncomingMessage struct {
	User  string
	Text  string
	Photo string
}

type IncomingMessageChan chan *IncomingMessage
