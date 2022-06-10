package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	// "math/rand"

	"os"
	"strings"

	"github.com/ArtyomArtamonov/wordl/internal/game"
	"github.com/ArtyomArtamonov/wordl/internal/util"
)

const FILENAME = "../english-nouns.txt"

var words []string

func main() {
	rand.Seed(time.Now().UnixNano())
	bytes, err := ioutil.ReadFile(FILENAME)
	util.FatalOnError(err)

	words = strings.Split(string(bytes), "\n")
	stdin := bufio.NewReader(os.Stdin)
	exit := "n"
	for strings.Trim(exit, "\n") != "y" {
		field := game.NewEmptyField()
		game := game.NewGame(field, stdin, &game.Wordlist{Words: words})
		word := words[rand.Intn(len(words))]

		game.Play(word)

		fmt.Println("Do you want to exit? (y/n)")
		exit, err = stdin.ReadString('\n')
		util.FatalOnError(err)
	}
}
