package slackrtm

import (
    "github.com/nlopes/slack"
)

func CreateSlackClient(apiKey string) *slack.RTM {
    api := slack.New(apiKey)
    rtm := api.NewRTM()
    go rtm.ManageConnection() 
    return rtm
}
