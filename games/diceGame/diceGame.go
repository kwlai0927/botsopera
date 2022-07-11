package dicegame

import "telegarm-bot-go/app"

type DiceGame struct {
}

type DiceGameBet struct {
	aaa    string
	amount int64
}

func isBetText(s string) (*DiceGameBet, error) {
	return &DiceGameBet{
		aaa:    "x",
		amount: 100,
	}, nil
}

func (g *DiceGame) HandleCusomerMessage(room *app.Room, username string, text string) {
	bet, err := isBetText(text)
	if err != nil {
		room.Table.UserBet(username, bet.amount)
	}
}
