package interactor

import (
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    // "github.com/MakeFang/GoUtility/sqldb"
)

type ReturnRes struct {
    Msg string
    Err error
    Code uint
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

var returnHelp ReturnRes = ReturnRes{Msg: helpString, Err: nil, Code: 0}

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

// func GetProcessing()
//
// func SetProcessing()

// var db *gorm.DB = sqldb.SetupDB()

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
