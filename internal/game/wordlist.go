package game

type Wordlist struct {
	Words []string
}

func (w *Wordlist) Contains(word string) bool {
	for _, v := range w.Words {
		if v == word {
			return true
		}
	}
	return false
}
