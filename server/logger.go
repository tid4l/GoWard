package server

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogFileName = timestamp() + "_GoWard.log"
	LogFile, _  = os.OpenFile(LogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Logger      = log.New(LogFile, "GW ", log.LstdFlags)
)

// Formats timestamp to append to log name upon log creation.
func timestamp() string {
	var timeFormatted string
	var month string
	var day string
	t := time.Now()
	year := t.Year()
	mm := t.Month()
	dd := t.Day()
	var m int = int(mm)
	if len(fmt.Sprint(m)) < 2 {
		month = "0" + fmt.Sprint(m)
	} else {
		month = fmt.Sprint(m)
	}
	if len(fmt.Sprint(dd)) < 2 {
		day = "0" + fmt.Sprint(dd)
	} else {
		day = fmt.Sprint(dd)
	}
	timeFormatted = fmt.Sprint(year) + month + day
	return timeFormatted
}
