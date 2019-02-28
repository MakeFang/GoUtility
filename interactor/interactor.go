package interactor

import (
    "time"
)

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
