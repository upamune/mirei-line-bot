package controllers

import (
	"github.com/labstack/echo/engine/standard"
	"google.golang.org/appengine/urlfetch"
	"net/http"

	"github.com/labstack/echo"
	"github.com/upamune/mirei-line-bot/models"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"io/ioutil"
	"net/http/httputil"
	"sync"
)

func LineCallBackHandler(c echo.Context) error {
	request := extractHttpRequestFromEchoContext(c)
	ctx := appengine.NewContext(request)
	req := &models.LineReceivingMessagesRequest{}
	err := c.Bind(req)
	if err != nil {
		dump, _ := httputil.DumpRequest(request, true)
		log.Errorf(ctx, "Bind Error: %v, %s, %s", err, string(dump))

		return c.String(http.StatusInternalServerError, "Invalid JSON")
	}
	log.Infof(ctx, "Decode Json Successfully %v", req)

	client := urlfetch.Client(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(len(req.Result))
	for _, result := range req.Result {
		go func(result *models.Result) {
			res := &models.Response{
				To:        []string{result.Content.From},
				ToChannel: 1383378250,
				EventType: "138311608800106203",
				Content: models.ResponseContent{
					ContentType: 1,
					ToType:      1,
					Text:        result.Content.Text + "ぷり",
				},
			}
			resp, err := res.Do(client)
			defer resp.Body.Close()
			if err != nil {
				log.Errorf(ctx, "Response Error: %v", err)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Errorf(ctx, "Read Body Error: %v", err)
				return
			}
			log.Infof(ctx, "SendResponse: %v", string(body))
			wg.Done()
		}(result)
	}
	wg.Wait()
	return c.String(http.StatusOK, "{}")
}

func extractHttpRequestFromEchoContext(c echo.Context) *http.Request {
	return c.Request().(*standard.Request).Request
}
