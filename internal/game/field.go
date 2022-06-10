package game

import (
	"regexp"
	"strings"

	"github.com/buger/goterm"
)

const NUM_OF_GUESSES = 6

const (
	HitNotHit = iota
	HitInWord
	HitInPosition
)

var isMatchingRegexp = regexp.MustCompile(`[A-Za-z]{5}`).MatchString

type Field struct {
	currentIndex  int
	field         []string
	unusedLetters string
}

func NewEmptyField() *Field {
	var res []string
	for i := 0; i < NUM_OF_GUESSES; i++ {
		res = append(res, "     ")
	}
	return &Field{
		currentIndex:  0,
		field:         res,
		unusedLetters: "q w e r t y u i o p  a s d f g h j k l  z x c v b n m",
	}
}

func (f *Field) push(word string) {
	if !isMatchingRegexp(word) {
		return
	}
	f.field[f.currentIndex] = strings.ToLower(word)

	f.currentIndex++
}

// renderWindow gets a word that is user is trying to guess
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
			if []rune(word)[i] == v {
				hit = HitInPosition
			}
			switch hit {
			case HitNotHit:
				if string(v) != " " {
					f.unusedLetters = strings.Replace(f.unusedLetters, string(v), "-", 1)
				}
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

	f.renderKeyboard()

	goterm.Flush()
}

func (f *Field) renderKeyboard() {
	// Split by two spaces in order to split string into three rows (like in a keyboard)
	lines := strings.Split(f.unusedLetters, "  ")

	goterm.MoveCursor(20, 4)
	goterm.Print(lines[0])

	goterm.MoveCursor(21, 5)
	goterm.Print(lines[1])

	goterm.MoveCursor(22, 6)
	goterm.Print(lines[2])

	goterm.MoveCursor(1, 11)
}
