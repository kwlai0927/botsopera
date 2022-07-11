package demoDrama

import (
	"time"
)

type DemoGame struct {
	needStop bool
}

func (g *DemoGame) start(chS chan<- string, chR <-chan string) {

	for {
		chS <- "開始投注"
		var bets []string
		tTen := time.After(20)
		tEnd := time.After(30)
	forbet:
		for {
			select {
			case bet := <-chR:
				bets = append(bets, bet)
			case <-tTen:
				chS <- "最後投注倒數 10 秒"
			case <-tEnd:
				break forbet
			}
		}
		chS <- "截止投注"
		time.Sleep(1)
		chS <- "截止投注"
		chS <- "顯示投注 png"
		chS <- "核對表格, 請仔細查看你的注單是否有,  若有誤，請您及時群內反饋, 以便核實修正，謝謝!"
		chS <- "開: 等待遊戲影像回傳中..."
		chS <- "顯示遊戲結果錄影"
		chS <- "顯示遊戲結果 png"
		chS <- "顯示玩家投注遊戲結果 png"
		if g.needStop {
			break
		}
	}

}

func (g *DemoGame) stopGame() {
	g.needStop = true
}
