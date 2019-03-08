package controller

import (
	interactor "github.com/MakeFang/GoUtility/interactor"
	"strings"
)

// ReturnRes is the return response.
type ReturnRes = interactor.ReturnRes

var helpString string = `Type in <operation> <arg1> <arg2> ...
  - get
  - set [time]
      - [time]:
          time in the format yyyy-mm-ddThh:mm:ss-08:00
  - cancel [reservationID]`

var returnHelp ReturnRes = ReturnRes{Msg: helpString, Err: nil}

// FormatCommands is a function for splitting input by spaces.
func FormatCommands(input string) []string {
	result := strings.Split(input, " ")
	// fmt.Println(result)
	return result
}

// ControllerRouting is a function that sends different command to diff func.
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
