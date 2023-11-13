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

func ConvertorLogPrintln(details ...string) {
	log.Println(strings.Join(details, " "))
}

func ConvertorLogPrintf(err error, details ...string) {
	log.Println(fmt.Sprintf("%s %+v",
		strings.Join(details, " "), err))
}
