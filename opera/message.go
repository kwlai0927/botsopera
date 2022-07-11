package opera

type MessageID string

type MessageImage string
type MessageVideo string

type Message struct {
	ID       MessageID
	UserName string
	Text     string
	Images   []MessageImage
	Videos   []MessageVideo
}

type BotMessage struct {
	RoomID  RoomID
	Message *Message
}
