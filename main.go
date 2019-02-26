package main

import (
    "fmt"
    // "log"
    "os"
    "strings"
    "github.com/nlopes/slack"
    _ "github.com/joho/godotenv/autoload"
)

func main() {
    api := slack.New(os.Getenv("BOT_OAUTH_ACCESS_TOKEN"))
    // logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
    // slack.SetLogger(logger)
    // api.SetDebug(true)

    rtm := api.NewRTM()
    go rtm.ManageConnection()

    for msg := range rtm.IncomingEvents {
        fmt.Println("Event Received: ", msg.Type)
        switch ev := msg.Data.(type) {
        case *slack.MessageEvent:
            fmt.Printf("%+v\n", ev.Msg)
            direct := strings.HasPrefix(ev.Msg.Channel, "D")
            if !direct || ev.Msg.SubType == "message_deleted" {
                fmt.Println("message not direct: ignored.")
                continue
            }
            fmt.Printf("%+v\n", rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
            rtm.SendMessage(rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
        }
    }
}
