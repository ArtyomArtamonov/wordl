package util

import (
	"log"

	"github.com/buger/goterm"
)

const INPUT_LINE_NUMBER = 13

func PrintStats(lineNumber int, info string, args ...interface{}) {
	PrintOn(lineNumber, 20, info, args...)
}

func PrintOn(lineNumber, columnNumber int, info string, args ...interface{}) {
	goterm.MoveCursor(columnNumber, lineNumber)
	goterm.Printf(info, args...)
	goterm.MoveCursor(1, INPUT_LINE_NUMBER)
	goterm.Flush()
}

func ClearInputLine() {
	goterm.MoveCursor(1, INPUT_LINE_NUMBER)
	goterm.Print(`                                                  `)
	goterm.MoveCursor(1, INPUT_LINE_NUMBER)
	goterm.Flush()
}

func FatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
