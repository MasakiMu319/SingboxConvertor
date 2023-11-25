package utils

import (
	"fmt"
	"log"
	"strings"
)

const (
	logHeader = "[Convertor] "
)

func init() {
	log.SetPrefix(logHeader)
}

// ConvertorLogPrintln help to print normal info.
func ConvertorLogPrintln(details ...string) {
	log.Println(strings.Join(details, " "))
}

// ConvertorLogPrintf help to print error info.
func ConvertorLogPrintf(err error, details ...string) {
	log.Println(fmt.Sprintf("%s %+v",
		strings.Join(details, " "), err))
}
