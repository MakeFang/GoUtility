[![Go Report Card](https://goreportcard.com/badge/github.com/MakeFang/GoUtility)](https://goreportcard.com/badge/github.com/MakeFang/GoUtility)

# Reserve Room Resource

## Usage

Start a direct message with MakeTime bot.

Type in [operation] [arg1] [arg2] ...
- get

The get command currently list all reservations

- set [time]
    - [time]: time in the format yyyy-mm-ddThh:mm:ss-08:00
    
Set reservation with starting time [time]. It books 30 minutes automatically right now.

- cancel [reservationID]

Cancel the user's reservation by id
