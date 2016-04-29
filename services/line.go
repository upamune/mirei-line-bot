package services

import (
	"github.com/upamune/line-bot-sdk-go/linebot"
	"github.com/upamune/mirei-line-bot/models"
	"sync"
	"net/http"
)

var once sync.Once
var bot *linebot.Client

func GetLineClient(clinet *http.Client) (*linebot.Client) {
	once.Do(func(){
		bot, _ = linebot.NewClient(models.GetLineChannelID(), models.GetLineChannelSecret(), models.GetLineChannelMID())
	})
	bot.SetHTTPClient(clinet)
	return bot

}
