package game

import (
	"regexp"
	"strings"

	"github.com/buger/goterm"
)

const (
	HitNotHit = iota
	HitInWord
	HitInPosition
)

var isMatchingRegexp = regexp.MustCompile(`[A-Za-z]{5}`).MatchString

type Field struct {
	currentIndex int
	field        []string
}

func NewEmptyField() *Field {
	var res []string
	for i := 0; i < 5; i++ {
		res = append(res, "     ")
	}
	return &Field{
		currentIndex: 0,
		field:        res,
	}
}

func (f *Field) push(word string) {
	if !isMatchingRegexp(word) {
		return
	}
	f.field[f.currentIndex] = strings.ToLower(word)

	f.currentIndex++
}

func (f *Field) renderWindow(word string) {
	goterm.Clear()
	goterm.MoveCursor(1, 1)
	for guess := 0; guess < len(f.field); guess++ {
		hittedInPositionIndexes := []int{}
		goterm.Print("|")
		// render letters
		for i, v := range f.field[guess] {
			hit := HitNotHit
			if strings.Contains(word, string(v)) {
				hit = HitInWord
			}
			if strings.IndexRune(word, v) == i {
				hit = HitInPosition
			}
			switch hit {
			case HitNotHit:
				goterm.Print(string(v) + "|")
			case HitInWord:
				goterm.Print(strings.ToUpper(string(v)) + "|")
			default:
				hittedInPositionIndexes = append(hittedInPositionIndexes, i)
				goterm.Print(strings.ToUpper(string(v)) + "|")
			}
		}
		goterm.Println()
		goterm.Print("|")
		// render dashes
		for i := 0; i < 5; i++ {
			toPrint := " "
			for _, v := range hittedInPositionIndexes {
				if v == i {
					toPrint = "-"
					break
				}
			}
			goterm.Print(toPrint + "|")
		}
		goterm.Println()
	}
	goterm.Println()
	goterm.Flush()
}
