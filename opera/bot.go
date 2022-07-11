package opera

type BotID string

// 遊戲處理流程，依據各種遊戲，有不同的處理

type Bot interface {
	Watch(ch chan *BotMessage)
	Register(chSend <-chan *BotMessage, chReceive chan<- *BotMessage)
}
