package interactor

import (
    "fmt"
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    // "github.com/MakeFang/GoUtility/sqldb"
)

type ReturnRes struct {
    Msg string
    Err error
    // Code uint
}

// type DBInterface interface {
//
// }

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

var helpString string = `Type in <operation> <options>
  - get
  - set
    [time]: time in the format yyyy-mm-ddThh:mm:ss-08:00`

var returnHelp ReturnRes = ReturnRes{Msg: helpString, Err: nil}

// #TODO: investigate other db options
var DB *gorm.DB

func SetDB(db *gorm.DB) {
    DB = db
    DB.AutoMigrate(&Reservation{}, &User{})
}

// type Rules struct {
//     NumBlocks int
//
// }

func GetProcessing() []Reservation {
    var reservations []Reservation
    DB.Find(&reservations)
    // response := fmt.Sprintf("%v", reservations)
    return reservations
}

func SetProcessing(t1 time.Time, userID string) ReturnRes {
    var reservation Reservation
    var resMsg string
    newRes := Reservation{StartTime: t1, UserSlackID: userID, RoomID: "1"}
    DB.Where(Reservation{StartTime: t1, UserSlackID: "", RoomID: "1"}).First(&reservation)
    fmt.Println("look up reservations", reservation)
    if reservation == (Reservation{}) {
        fmt.Println("no records")
        DB.Create(&newRes)
        // db.Where("StartTime = ?", t1).First(&reservation)
        // fmt.Println(reservation)
        resMsg = "Success. Reserved " + t1.Format(time.RFC3339)
        return ReturnRes{Msg: resMsg, Err: nil}
    }
    // } else {
    //     resMsg = "Failure. Time slot taken already."
    // }
    resMsg = "Failure. Time slot taken already."
    return ReturnRes{Msg: resMsg, Err: nil}
}

// var db *gorm.DB = sqldb.SetupDB()

func GetParsing(args []string, userID string) ReturnRes {
    reservations := GetProcessing()
    response := fmt.Sprintf("%v", reservations)
    return ReturnRes{Msg: response, Err: nil}
    // numArgs := len(args)
    // if numArgs != 1 {
    //     return "help", nil
    // }
    // t1, e := time.Parse(time.RFC3339, args[0])
    // if e != nil {
    //     return "help", e
    // }
    // startTime := t1.Truncate(30*time.Minute)
    // return startTime.Format(time.RFC3339), nil
}

func SetParsing(args []string, userID string) ReturnRes {
    numArgs := len(args)
    if numArgs != 1 {
        return returnHelp
    }
    t1, e := time.Parse(time.RFC3339, args[0])
    if e != nil {
        return ReturnRes{Msg: helpString, Err: e}
    }
    startTime := t1.Truncate(30*time.Minute)
    response := SetProcessing(startTime, userID)
    return response
}
