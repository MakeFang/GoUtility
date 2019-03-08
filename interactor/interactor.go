package interactor

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// ReturnRes is a struct with returning message string and err.
type ReturnRes struct {
	Msg string
	Err error
	// Code uint
}

// Reservation is a struct schema with details about the resevation.
type Reservation struct {
	gorm.Model
	StartTime   time.Time
	UserSlackID string
	RoomID      string
}

// User is a struct schema for keeping track of users reservations
type User struct {
	gorm.Model
	SlackID      string
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

// DB is the gorm database.
var DB *gorm.DB

// SetDB is a function that sets value to DB variable.
func SetDB(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&Reservation{}, &User{})
}

// GetProcessing is function that process get requests
func GetProcessing() []Reservation {
	var reservations []Reservation
	DB.Find(&reservations)
	// response := fmt.Sprintf("%v", reservations)
	return reservations
}

// SetProcessing is function that process set requests
func SetProcessing(t1 time.Time, userID string) ReturnRes {
	var reservation Reservation
	var resMsg string
	newRes := Reservation{StartTime: t1, UserSlackID: userID, RoomID: "1"}
	DB.Where(Reservation{StartTime: t1, UserSlackID: "", RoomID: "1"}).First(&reservation)
	fmt.Println("look up reservations", reservation)
	if reservation == (Reservation{}) {
		fmt.Println("no records")
		DB.Create(&newRes)
		fmt.Printf("%+v\n", newRes)
		// db.Where("StartTime = ?", t1).First(&reservation)
		// fmt.Println(reservation)
		resMsg = "Success. Reserved " + t1.Format(time.RFC3339)
		return ReturnRes{Msg: resMsg, Err: nil}
	}
	resMsg = "Failure. Time slot taken already."
	return ReturnRes{Msg: resMsg, Err: nil}
}

// CancelProcessing is function that process cancel requests
func CancelProcessing(reservationID string, userID string) ReturnRes {
	// var reservations []Reservation
	var resMsg string
	var toDelete Reservation
	DB.First(&toDelete, reservationID)
	if toDelete.UserSlackID == userID {
		resMsg = "Reservations canceled"
		DB.Delete(&toDelete)
		return ReturnRes{Msg: resMsg, Err: nil}
	}
	// DB.Where(Reservation{StartTime: time.Time{}, UserSlackID: userID , RoomID: "1"}).Delete(Reservation{})
	resMsg = "Slack ID mismatch. Unable to cancel reservation."
	return ReturnRes{Msg: resMsg, Err: nil}
}

// var db *gorm.DB = sqldb.SetupDB()

// GetParsing is function that parse get requests
func GetParsing(args []string, userID string) ReturnRes {
	// valid := GetVal(args)
	reservations := GetProcessing()
	response := fmt.Sprintf("%v", reservations)
	return ReturnRes{Msg: response, Err: nil}
}

// SetParsing is function that parse set requests
func SetParsing(args []string, userID string) ReturnRes {
	numArgs := len(args)
	if numArgs != 1 {
		return returnHelp
	}
	t1, e := time.Parse(time.RFC3339, args[0])
	if e != nil {
		return ReturnRes{Msg: helpString, Err: e}
	}
	startTime := t1.Truncate(30 * time.Minute)
	response := SetProcessing(startTime, userID)
	return response
}

// CancelParsing is function that parse cancel requests
func CancelParsing(args []string, userID string) ReturnRes {
	numArgs := len(args)
	if numArgs != 1 {
		return returnHelp
	}
	reservationID := args[0]
	response := CancelProcessing(reservationID, userID)
	return response
}
