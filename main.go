package main

import (
    "fmt"
    // "log"
    "os"
    "strings"
    "github.com/nlopes/slack"
    _ "github.com/joho/godotenv/autoload"
)

func FormatCommands(input string) []string {
    result := strings.Split(input, " ")
    // fmt.Println(result)
    return result
}

// controller
func ControllerRouting(args []string) string {
    numArgs := len(args)
    if numArgs > 3 || numArgs < 1 {
        return "size bad"
    }
    switch firstArg := args[0]; firstArg {
    case "get":
        return "get"
    case "set":
        return "set"
    default:
        return "help"
    }
    // for _, arg := range args {
    //
    // }
}

// schema

// Private logic

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
            // fmt.Printf("%+v\n", rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
            // fmt.Println(FormatCommands(ev.Msg.Text))
            formattedMsg := FormatCommands(ev.Msg.Text)
            outputMsg := ControllerRouting(formattedMsg)
            rtm.SendMessage(rtm.NewOutgoingMessage(outputMsg, ev.Msg.Channel))
        }
    }
}
