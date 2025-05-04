package common

type Message struct {
	ChatId     int64
	MessagesId []int
	Response   bool
}

var MessageInProgress *Message
