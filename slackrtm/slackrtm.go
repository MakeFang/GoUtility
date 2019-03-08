package slackrtm

import (
    "fmt"
    "strings"
    "github.com/nlopes/slack"
    "github.com/MakeFang/GoUtility/controller"
)

// CreateSlackClient is a function that initiate the slack rtm.
func CreateSlackClient(apiKey string) *slack.RTM {
    api := slack.New(apiKey)
    rtm := api.NewRTM()
    go rtm.ManageConnection()
    return rtm
}

// GetIncomingMsg is a function that deals with incoming funciton.
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
            // var user User
            // var allUsers []User
            // // fmt.Println(ev.Msg.User)
            // // fmt.Printf("%+v\n", User{SlackID: ev.Msg.User})
            // db.Where(User{SlackID: ev.Msg.User}).FirstOrCreate(&user)
            // db.Find(&allUsers)
            // fmt.Println(allUsers)
            // fmt.Println(user)
            // // db.Where("SlackID = ?", ev.Msg.User).First(&user)
            // // fmt.Printf("%+v\n", user)
            // // fmt.Printf("%+v\n", rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
            // // fmt.Println(FormatCommands(ev.Msg.Text))
            formattedMsg := controller.FormatCommands(ev.Msg.Text)
            // fmt.Println(formattedMsg)
            outputMsg := controller.ControllerRouting(formattedMsg, ev.Msg.User)
            // fmt.Println(outputMsg.Msg)
            slackClient.SendMessage(slackClient.NewOutgoingMessage(outputMsg.Msg, ev.Msg.Channel))
        }
    }
}
