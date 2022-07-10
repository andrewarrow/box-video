package audio

import (
	"io/ioutil"
	"strings"
)

var wordLines [][]*Word
var wordLineIndex = 0
var wordChange = false
var wordIndex = 0
var wordChars = 0
var words []*Word

type Word struct {
	Word string
	Time int
}

func ReadWordTimes(filename string) {
	b, _ := ioutil.ReadFile(filename)

	wordLines = [][]*Word{}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}

		words := strings.Split(trimmed, " ")

		wordLine := []*Word{}
		for _, word := range words {
			w := Word{word, 1000}
			wordLine = append(wordLine, &w)
		}
		wordLines = append(wordLines, wordLine)
	}
}
