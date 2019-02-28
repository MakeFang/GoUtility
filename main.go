package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "time"
    "github.com/MakeFang/GoUtility/slackrtm"
    "github.com/nlopes/slack"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    _ "github.com/joho/godotenv/autoload"
)

func FormatCommands(input string) []string {
    result := strings.Split(input, " ")
    // fmt.Println(result)
    return result
}

// controller
func ControllerRouting(args []string, db *gorm.DB, userID string) string {
    numArgs := len(args)
    if numArgs > 3 || numArgs < 1 {
        return "help"
    }
    switch firstArg := args[0]; firstArg {
    case "get":
        response, err := GetParsing(args[1:])
        if err != nil {
            log.Fatal(err)
        }
        var reservations []Reservation
        db.Find(&reservations)
        response = fmt.Sprintf("%v", reservations)
        return response
    case "set":
        response, err := SetParsing(args[1:])
        if err != nil {
            log.Fatal(err)
        }
        t1, e := time.Parse(time.RFC3339, response)
        if e != nil {
            return "help"
        }
        var reservation Reservation
        var resMsg string
        newRes := Reservation{StartTime: t1, UserSlackID: userID, RoomID: "1"}
        db.Where(Reservation{StartTime: t1, UserSlackID: "", RoomID: "1"}).First(&reservation)
        fmt.Println("look up reservations", reservation)
        if reservation == (Reservation{}) {
            fmt.Println("no records")
            db.Create(&newRes)
            // db.Where("StartTime = ?", t1).First(&reservation)
            // fmt.Println(reservation)
            resMsg = "Success. Reserved " + t1.Format(time.RFC3339)
        } else {
            resMsg = "Failure. Time slot taken already."
        }
        return resMsg
    default:
        return "help"
    }
    // for _, arg := range args {
    //
    // }
}

// #TODO: disallow booking time that has passed

func GetParsing(args []string) (string, error) {
    numArgs := len(args)
    if numArgs != 1 {
        return "help", nil
    }
    t1, e := time.Parse(time.RFC3339, args[0])
    if e != nil {
        return "help", e
    }
    startTime := t1.Truncate(30*time.Minute)
    return startTime.Format(time.RFC3339), nil
}

func SetParsing(args []string) (string, error) {
    numArgs := len(args)
    if numArgs != 1 {
        return "help", nil
    }
    t1, e := time.Parse(time.RFC3339, args[0])
    if e != nil {
        return "help", e
    }
    startTime := t1.Truncate(30*time.Minute)
    return startTime.Format(time.RFC3339), nil
}

// schema
type Reservation struct {
    gorm.Model
    StartTime time.Time
    UserSlackID string
    RoomID string
}

type User struct {
    gorm.Model
    SlackID string
    Reservations []Reservation `gorm:"foreignkey:UserSlackID;association_foreignkey:SlackID"`
}

// Private logic

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect to database")
    }
    defer db.Close()

    db.AutoMigrate(&Reservation{}, &User{})

    botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
    slackClient := slackrtm.CreateSlackClient(botToken)
    slack.GetIncomingMsg(slackClient)

    // api := slack.New(botToken)
    // logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
    // slack.SetLogger(logger)
    // api.SetDebug(true)
    //
    // rtm := api.NewRTM()
    // go rtm.ManageConnection()

    // for msg := range slackClient.IncomingEvents {
    //     fmt.Println("Event Received: ", msg.Type)
    //     switch ev := msg.Data.(type) {
    //     case *slack.MessageEvent:
    //         fmt.Printf("%+v\n", ev.Msg)
    //         direct := strings.HasPrefix(ev.Msg.Channel, "D")
    //         if !direct || ev.Msg.SubType == "message_deleted" {
    //             fmt.Println("message not direct: ignored.")
    //             continue
    //         }
    //         var user User
    //         var allUsers []User
    //         // fmt.Println(ev.Msg.User)
    //         // fmt.Printf("%+v\n", User{SlackID: ev.Msg.User})
    //         db.Where(User{SlackID: ev.Msg.User}).FirstOrCreate(&user)
    //         db.Find(&allUsers)
    //         fmt.Println(allUsers)
    //         fmt.Println(user)
    //         // db.Where("SlackID = ?", ev.Msg.User).First(&user)
    //         // fmt.Printf("%+v\n", user)
    //         // fmt.Printf("%+v\n", rtm.NewOutgoingMessage("hello", ev.Msg.Channel))
    //         // fmt.Println(FormatCommands(ev.Msg.Text))
    //         formattedMsg := FormatCommands(ev.Msg.Text)
    //         outputMsg := ControllerRouting(formattedMsg, db, ev.Msg.User)
    //         slackClient.SendMessage(slackClient.NewOutgoingMessage(outputMsg, ev.Msg.Channel))
    //     }
    // }
}
