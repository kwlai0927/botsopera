package opera

type DramaID string
type DramaName string

// 遊戲處理流程，依據各種遊戲，有不同的處理

type Drama interface {
	ID() DramaID
	Name() DramaName
	Register(chSend <-chan *Message, chReceive chan<- *Message)
}
