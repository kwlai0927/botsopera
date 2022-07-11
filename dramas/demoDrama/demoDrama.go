package demoDrama

import (
	"github.com/kwlai0927/botsopera/opera"
)

type DemoDrama struct{}

func (d *DemoDrama) Register(chSend chan<- *opera.Message, chReceive <-chan *opera.Message) {
	game := &DemoGame{
		needStop: false,
	}
	chGS := make(chan string)
	chGR := make(chan string)
	go game.start(chGR, chGS)
loop:
	for {
		select {
		case txt := <-chGR:
			chSend <- &opera.Message{
				Text: txt,
			}
		case msg, ok := <-chReceive:
			if !ok {
				game.stopGame()
				close(chSend)
				close(chGR)
				break loop
			}
			chGS <- msg.Text
		}
	}
}
