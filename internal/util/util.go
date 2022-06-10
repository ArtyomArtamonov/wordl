package util

import (
	"log"

	"github.com/buger/goterm"
)

func PrintStats(lineNumber int, info string, args ...interface{}) {
	goterm.MoveCursor(20, lineNumber)
	goterm.Printf(info, args...)
	goterm.MoveCursor(1, 11)
	goterm.Flush()
}

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
