package controllers

import (
	"github.com/labstack/echo/engine/standard"
	"google.golang.org/appengine/urlfetch"
	"net/http"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/upamune/mirei-line-bot/services"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"sync"
)

func LineCallBackHandler(c echo.Context) error {
	request := extractHttpRequestFromEchoContext(c)
	ctx := appengine.NewContext(request)
	client := urlfetch.Client(ctx)

	bot, err := services.GetLineClient(client)
	if err != nil {
		log.Errorf(ctx, "Line Clinet Error: %v", err)
	}
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
			if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
				// すでにifでチェックしているので，ここのエラーは握り潰しても良い
				textContent, _ := content.TextContent()
				res, err := bot.SendText([]string{content.From}, textContent.Text+"ぷり")
				if err != nil {
					log.Errorf(ctx, "Send Message Text Error: %v", err)
					return
				}
				log.Infof(ctx, "Succeed Sent Messages: %v", res)
			}
		}(&result)
	}
	wg.Wait()
	return c.String(http.StatusOK, "{}")
}

func extractHttpRequestFromEchoContext(c echo.Context) *http.Request {
	return c.Request().(*standard.Request).Request
}
