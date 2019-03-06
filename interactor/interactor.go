package interactor

import (
    "fmt"
    "time"
    // "errors"
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

var helpString string = `Type in <operation> <arg1> <arg2> ...
  - get [roomID]
      - [roomID]:
          (optional) 1 for room#1 or 2 for room#2.
          If roomID not provided, will list all reservations
          for the user.
  - set [roomID] [time]
      - [roomID]:
          (NOT OPTIONAL) 1 for room#1 or 2 for room#2.
      - [time]:
          (NOT OPTIONAL) time in the format yyyy-mm-ddThh:mm:ss-08:00
  - cancel [reservationID]`

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

// func InputValidation() {
//
// }

// type GetArgs struct {
//     roomID string
// }
//
// func GetVal(args []string) (GetArgs, error){
//     if len(args) > 1 {
//         return GetArgs{}, errors.NEW("")
//     }
// }
//
// func SetVal() bool{
//
// }
//
// func CancelVal() bool{
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
        fmt.Printf("%+v\n",newRes)
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
    // valid := GetVal(args)
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

func CancelParsing(args []string, userID string) ReturnRes {
    return ReturnRes{Msg: "cancel", Err: nil}
}
