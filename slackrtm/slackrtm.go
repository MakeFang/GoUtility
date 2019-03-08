package slackrtm

import (
	"fmt"
	"github.com/MakeFang/GoUtility/controller"
	"github.com/nlopes/slack"
	"strings"
)

// CreateSlackClient is a function that initiate the slack rtm.
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	return rtm
}

// GetIncomingMsg is a function that deals with incoming function.
func GetIncomingMsg(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("%+v\n", ev.Msg)
			direct := strings.HasPrefix(ev.Msg.Channel, "D")
			if !direct || ev.Msg.SubType == "message_deleted" {
				fmt.Println("message not direct: ignored.")
				continue
			}
			formattedMsg := controller.FormatCommands(ev.Msg.Text)
			outputMsg := controller.ControllerRouting(formattedMsg, ev.Msg.User)
			slackClient.SendMessage(slackClient.NewOutgoingMessage(outputMsg.Msg, ev.Msg.Channel))
		}
	}
}
