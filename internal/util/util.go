package util

import (
	"log"

	"github.com/buger/goterm"
)

func PrintStats(lineNumber int, info string, args ...interface{}) {
	PrintOn(lineNumber, 20, info, args...)
}

func PrintOn(lineNumber, columnNumber int, info string, args ...interface{}) {
	goterm.MoveCursor(columnNumber, lineNumber)
	goterm.Printf(info, args...)
	goterm.MoveCursor(1, 11)
	goterm.Flush()
}

func ClearInputLine() {
	goterm.MoveCursor(1, 11)
	goterm.Print(`                                                  `)
	goterm.MoveCursor(1, 11)
	goterm.Flush()
}

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
