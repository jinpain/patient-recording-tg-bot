package common

type Message struct {
	ChatId     int64
	MessagesId []int
}

var MessageInProgress *Message
