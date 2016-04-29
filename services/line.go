package services

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/upamune/mirei-line-bot/models"
	"net/http"
)

func GetLineClient(client *http.Client) (*linebot.Client, error) {
	bot, err := linebot.NewClient(models.GetLineChannelID(), models.GetLineChannelSecret(), models.GetLineChannelMID(), linebot.WithHTTPClient(client))
	return bot, err
}
