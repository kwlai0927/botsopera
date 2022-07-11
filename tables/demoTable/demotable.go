package demoTable

import (
	"fmt"
	"log"
	"strconv"
	"telegarm-bot-go/app"
	"time"

	"github.com/spf13/viper"
)

type DemoGame struct {
}

func (g *DemoGame) HandleCusomerMessage(room *app.Room, username string, text string) {

}

type DemoTable struct {
	game  *DemoGame
	rooms map[int64]*app.Room
}

func (t *DemoTable) GetGame() *DemoGame {
	return t.game
}

func (t *DemoTable) LinkGame() {

	for _, room := range t.rooms {

		// 開始遊戲
		room.SendMessage("請玩家開始投注")

		time.AfterFunc(time.Duration(gameLastBetTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_LAST_BET_TIMER")
			room.SendMessage("最後投注倒數 10 秒")
		})
		time.AfterFunc(time.Duration(gameStopBetTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_STOP_BET_TIMER")
			room.SendMessage("截止投注")
		})

		time.AfterFunc(time.Duration(gameShowBetImageTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_SHOW_BET_IMAGE_TIMER")
			room.SendMessage("顯示投注 png")
		})

		time.AfterFunc(time.Duration(gameCheckBetTableTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_CHECK_BET_TABLE_TIMER")
			room.SendMessage("核對表格, 請仔細查看你的注單是否有,  若有誤，請您及時群內反饋, 以便核實修正，謝謝!")
		})

		time.AfterFunc(time.Duration(gameWaitResultVideoTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_WAIT_RESULT_VIDEO_TIMER")
			room.SendMessage("開: 等待遊戲影像回傳中...")
		})

		time.AfterFunc(time.Duration(gameShowResultVideoTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_SHOW_RESULT_VIDEO_TIMER")
			room.SendMessage("顯示遊戲結果錄影")
		})

		time.AfterFunc(time.Duration(gameShowResultImageTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_SHOW_RESULT_IMAGE_TIMER")
			room.SendMessage("顯示遊戲結果 png")
		})

		time.AfterFunc(time.Duration(gameShowBetResultImageTimer)*time.Second, func() {
			fmt.Println("AfterFunc GAME_SHOW_BET_RESULT_IMAGE_TIMER")
			room.SendMessage("顯示玩家投注遊戲結果 png")
		})

	}

}

var (
	botChatID  int64 = 0
	gameChatID int64 = 0

	gameLastBetTimer            int64 = 0
	gameStopBetTimer            int64 = 0
	gameShowBetImageTimer       int64 = 0
	gameCheckBetTableTimer      int64 = 0
	gameWaitResultVideoTimer    int64 = 0
	gameShowResultVideoTimer    int64 = 0
	gameShowResultImageTimer    int64 = 0
	gameShowBetResultImageTimer int64 = 0
)

func GetViperData() {
	var err error
	botChatID, err = strconv.ParseInt(viper.GetString("bot_chat_id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("botChatID: ", botChatID)

	gameChatID, err = strconv.ParseInt(viper.GetString("game_chat_id"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameChatID: ", gameChatID)

	gameLastBetTimer, err := strconv.ParseInt(viper.GetString("game_last_bet_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameLastBetTimer: ", gameLastBetTimer)

	gameStopBetTimer, err = strconv.ParseInt(viper.GetString("game_stop_bet_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameStopBetTimer: ", gameStopBetTimer)

	gameShowBetImageTimer, err = strconv.ParseInt(viper.GetString("game_show_bet_image_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameShowBetImageTimer: ", gameShowBetImageTimer)

	gameCheckBetTableTimer, err = strconv.ParseInt(viper.GetString("game_check_bet_table_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameCheckBetTableTimer: ", gameCheckBetTableTimer)

	gameWaitResultVideoTimer, err = strconv.ParseInt(viper.GetString("game_wait_result_video_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameWaitResultVideoTimer: ", gameWaitResultVideoTimer)

	gameShowResultVideoTimer, err = strconv.ParseInt(viper.GetString("game_show_result_video_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameShowResultVideoTimer: ", gameShowResultVideoTimer)

	gameShowResultImageTimer, err = strconv.ParseInt(viper.GetString("game_show_result_image_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameShowResultImageTimer: ", gameShowResultImageTimer)

	gameShowBetResultImageTimer, err = strconv.ParseInt(viper.GetString("game_show_bet_result_image_timer"), 10, 64)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("gameShowBetResultImageTimer: ", gameShowBetResultImageTimer)
}
