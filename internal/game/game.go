package game

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/ArtyomArtamonov/wordl/internal/util"
)

type Game struct {
	field    *Field
	std      *bufio.Reader
	wordlist *Wordlist
}

func NewGame(field *Field, std *bufio.Reader, w *Wordlist) *Game {
	return &Game{
		field:    field,
		std:      std,
		wordlist: w,
	}
}

func (g *Game) Play(word string) {
	fmt.Print("Guess 5-letter word\n")
	field := NewEmptyField()
	guessesLeft := NUM_OF_GUESSES

	field.renderWindow(word)
	for {
		// util.PrintStats(8, "The word is %s\n", word) // Debug
		if guessesLeft == 0 {
			util.PrintStats(7, "You lost. The word was %s\n", word)
			break
		}
		util.PrintStats(2, "Guesses left: %d/%d\n", guessesLeft, NUM_OF_GUESSES)
		util.ClearInputLine()
		fmt.Print("Your guess: ")

		var guess string

		var err error
		guess, err = g.std.ReadString('\n')
		util.FatalOnError(err)
		guess = strings.Trim(strings.ToLower(guess), "\n")

		if !g.wordlist.Contains(guess) {
			util.PrintStats(3, "There is no such word in a wordlist, try again")
			continue
		}

		field.push(guess)
		guessesLeft -= 1

		field.renderWindow(word)
		if word == guess {
			util.PrintStats(7, "You won! The word is " + word)
			break
		}
	}
}
