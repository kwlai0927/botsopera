package telegramBot

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/kwlai0927/botsopera/opera"
)

type TelegramBot struct {
	botApi *tgbotapi.BotAPI
}

func (b *TelegramBot) Register(chSend chan<- *opera.BotMessage, chReceive <-chan *opera.BotMessage) {
	updateChan := b.botApi.GetUpdatesChan(tgbotapi.NewUpdate(0))
loop:
	for {
		select {
		case update := <-updateChan:
			if update.Message != nil {

				log.Println("Chat ID: ", update.Message.Chat.ID)
				log.Println("username: ", update.Message.From.UserName)
				log.Println("Text: ", update.Message.Text)

				chSend <- &opera.BotMessage{
					RoomID: opera.RoomID(strconv.FormatInt(update.Message.Chat.ID, 10)),
					Message: &opera.Message{
						ID:   "a message id",
						Text: update.Message.Text,
					},
				}
			}
		case msg, ok := <-chReceive:
			if !ok {
				break loop
			}
			tgChattable := makeTgChattableFromMessage(msg)
			b.botApi.Send(tgChattable)
		}
	}
}

func makeTgChattableFromMessage(msg *opera.BotMessage) tgbotapi.Chattable {
	chatID, _ := strconv.ParseInt(string(msg.RoomID), 10, 64)
	return tgbotapi.NewMessage(chatID, msg.Message.Text)
}
