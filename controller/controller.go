package controller

import (
    "strings"
    interactor "github.com/MakeFang/GoUtility/interactor"
)

// type returnRes struct {
//     Msg string
//     Err error
//     Code uint
// }
//

type ReturnRes = interactor.ReturnRes

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

func FormatCommands(input string) []string {
    result := strings.Split(input, " ")
    // fmt.Println(result)
    return result
}

func ControllerRouting(args []string, userID string) ReturnRes {
    numArgs := len(args)
    if numArgs > 3 || numArgs < 1 {
        return returnHelp
    }
    switch operation := args[0]; operation {
    case "get":
        response := interactor.GetParsing(args[1:], userID)
        return response
    case "set":
        response := interactor.SetParsing(args[1:], userID)
        return response
    case "cancel":
        response := interactor.CancelParsing(args[1:], userID)
        return response
    default:
        return returnHelp
    }
}

// func ControllerRouting(args []string, db *gorm.DB, userID string) string {
//     numArgs := len(args)
//     if numArgs > 3 || numArgs < 1 {
//         return "help"
//     }
//     switch firstArg := args[0]; firstArg {
//     case "get":
//         response, err := GetParsing(args[1:])
//         if err != nil {
//             log.Fatal(err)
//         }
//         var reservations []Reservation
//         db.Find(&reservations)
//         response = fmt.Sprintf("%v", reservations)
//         return response
//     case "set":
//         response, err := SetParsing(args[1:])
//         if err != nil {
//             log.Fatal(err)
//         }
//         t1, e := time.Parse(time.RFC3339, response)
//         if e != nil {
//             return "help"
//         }
//         var reservation Reservation
//         var resMsg string
//         newRes := Reservation{StartTime: t1, UserSlackID: userID, RoomID: "1"}
//         db.Where(Reservation{StartTime: t1, UserSlackID: "", RoomID: "1"}).First(&reservation)
//         fmt.Println("look up reservations", reservation)
//         if reservation == (Reservation{}) {
//             fmt.Println("no records")
//             db.Create(&newRes)
//             // db.Where("StartTime = ?", t1).First(&reservation)
//             // fmt.Println(reservation)
//             resMsg = "Success. Reserved " + t1.Format(time.RFC3339)
//         } else {
//             resMsg = "Failure. Time slot taken already."
//         }
//         return resMsg
//     default:
//         return "help"
//     }
//     // for _, arg := range args {
//     //
//     // }
// }
