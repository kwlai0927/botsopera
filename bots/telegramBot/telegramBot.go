package telegramBot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/kwlai0927/botsopera/opera"
)

// 處理Bot對接
type TelegramBot struct {
	botApi *tgbotapi.BotAPI
}

func (b *TelegramBot) Watch(ch chan *opera.Message) {
	updateChan := b.botApi.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updateChan {
		if update.Message != nil {

			log.Println("Chat ID: ", update.Message.Chat.ID)
			log.Println("username: ", update.Message.From.UserName)
			log.Println("Text: ", update.Message.Text)

			ch <- &opera.Message{
				ID:   "a message id",
				Text: update.Message.Text,
			}

		}
	}

}

func makeTgChattableFromMessage(msg *opera.Message) tgbotapi.Chattable {
	return tgbotapi.NewMessage(msg.ID, msg.Text)
}

func (b *TelegramBot) SendMessage(id opera.RoomID, msg *opera.Message) (*opera.BotMessage, error) {
	tgChattable := makeTgChattableFromMessage(msg)
	_, err := b.botApi.Send(tgChattable)
	if err != nil {
		return nil, err
	}
	return &opera.BotMessage{
		RoomID:  id,
		Message: msg,
	}, nil
}
