package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"

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
	rand.Seed(time.Now().UnixNano())

	fmt.Print("Guess 5-letter word\n")
	field := NewEmptyField()
	// unusedLetters := "abcdefghijklmnopqrstuvwxyz"
	guessesLeft := 5

	field.renderWindow(word)
	for {
		if guessesLeft == 0 {
			fmt.Printf("You lost. The word was %s\n", word)
			break
		}
		util.PrintStats(2, "Guesses left: %d/5\n", guessesLeft)
		var guess string

		fmt.Print("Your guess: ")
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
			fmt.Println("You won! The word is " + word)
			break
		}
	}
}
