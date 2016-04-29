package models

import (
	"os"
	"strconv"
)

const EndPoint = "https://trialbot-api.line.me/v1/events"

type LineReceivingMessagesRequest struct {
	Result []*Result `json:"result"`
}

type Result struct {
	From        string   `json:"from"`
	FromChannel int      `json:"fromChannel"`
	To          []string `json:"to"`
	ToChannel   int      `json:"toChannel"`
	EventType   string   `json:"eventType"`
	ID          string   `json:"id"`
	Content     Content  `json:"content"`
}

type Content struct {
	ToType          int             `json:"toType"`
	CreatedTime     int64           `json:"createdTime"`
	From            string          `json:"from"`
	Location        interface{}     `json:"location"`
	ID              string          `json:"id"`
	To              []string        `json:"to"`
	Text            string          `json:"text"`
	ContentMetaData ContentMetaData `json:"contentMetadata"`
	DeliveredTime   int             `json:"deliveredTime"`
	ContentType     int             `json:"contentType"`
	Seq             interface{}     `json:"seq"`
}

type ContentMetaData struct {
	ATRECVMODE     string `json:"AT_RECV_MODE"`
	SKIPBADGECOUNT string `json:"SKIP_BADGE_COUNT"`
}

type ResponseContent struct {
	ContentType int    `json:"contentType"`
	ToType      int    `json:"toType"`
	Text        string `json:"text"`
}

type Response struct {
	To        []string        `json:"to"`
	ToChannel int             `json:"toChannel"`
	EventType string          `json:"eventType"`
	Content   ResponseContent `json:"content"`
}

var linetChannelID = os.Getenv("LINE_CHANNEL_ID")
var linetChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
var linetChannelMID = os.Getenv("LINE_CHANNEL_MID")

func GetLineChannelID() int64 {
	i, err := strconv.ParseInt(linetChannelID, 10, 64)
	if err != nil {
		i = 0
	}
	return i
}

func GetLineChannelSecret() string {
	return linetChannelSecret
}

func GetLineChannelMID() string {
	return linetChannelMID
}
