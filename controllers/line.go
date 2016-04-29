package controllers

import (
	"google.golang.org/appengine/urlfetch"
	"net/http"

	"github.com/labstack/echo"
	"github.com/upamune/line-bot-sdk-go/linebot"
	"github.com/upamune/mirei-line-bot/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"sync"
	"github.com/labstack/echo/engine/standard"
)

func LineCallBackHandler(c echo.Context) error {
	request := extractHttpRequestFromEchoContext(c)
	ctx := appengine.NewContext(request)
	client := urlfetch.Client(ctx)

	bot := services.GetLineClient(client)
	received, err := bot.ParseRequest(request)
	if err != nil {
		log.Errorf(ctx, "Request Parse Error: %v", err)
		return c.String(http.StatusInternalServerError, "Request Parse Error")
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(received.Results))
	for _, result := range received.Results {
		go func(result *linebot.ReceivedResult) {
			defer wg.Done()
			content := result.Content()
			if content != nil && content.IsMessage {
				switch content.ContentType {
				case linebot.ContentTypeText:
					// すでにifでチェックしているので，ここのエラーは握り潰しても良い
					textContent, _ := content.TextContent()
					log.Infof(ctx, "Receive:Text:%s:%s", content.From, textContent.Text)
					res, err := bot.SendText([]string{content.From}, textContent.Text+"ぷり")
					if err != nil {
						log.Errorf(ctx, "Send Message Text Error: %v", err)
						return
					}
					log.Infof(ctx, "Succeed Sent Messages: %v", res)
				case linebot.ContentTypeImage:
					res, err := bot.SendText([]string{content.From}, "いい写真ぷり!")
					log.Infof(ctx, "Receive:Image:%s:%s", content.From)
					if err != nil {
						log.Errorf(ctx, "Send Message Text Error: %v", err)
						return
					}
					log.Infof(ctx, "Succeed Sent Messages: %v", res)
				case linebot.ContentTypeLocation:
					res, err := bot.SendText([]string{content.From}, "素敵な場所ぷりね")
					locationContent, _ := content.LocationContent()
					log.Infof(ctx, "Receive:Location:%s:%s", content.From, locationContent.Address)
					if err != nil {
						log.Errorf(ctx, "Send Message Text Error: %v", err)
						return
					}
					log.Infof(ctx, "Succeed Sent Messages: %v", res)
				case linebot.ContentTypeSticker:
					res, err := bot.SendText([]string{content.From}, "いいスタンプぷり! みれぃのスタンプも買って欲しいぷり💕")
					stickerContent, _ := content.StickerContent()
					log.Infof(ctx, "Receive:Sticker:%s:%d:%d:%d", content.From, stickerContent.PackageID, stickerContent.ID, stickerContent.Version)
					if err != nil {
						log.Errorf(ctx, "Send Message Text Error: %v", err)
						return
					}
					log.Infof(ctx, "Succeed Sent Messages: %v", res)
				}
			}
		}(&result)
	}
	wg.Wait()
	return c.String(http.StatusOK, "{}")
}

func extractHttpRequestFromEchoContext(c echo.Context) *http.Request {
	return c.Request().(*standard.Request).Request
}
