package main

import (
	"fmt"
	"math"
	"unicode"

	"example.com/lib50"
)

func readability(input string) int {
	var spaces, chars, periods int
	for _, c := range input {
		if unicode.IsSpace(c) {
			spaces++
		} else if c == '!' || c == '.' || c == '?' {
			periods++
		} else if unicode.IsLetter(c) {
			chars++
		}
	}

	var L float64
	var S float64

	L = float64(100*chars) / float64((spaces + 1))
	S = float64(100*periods) / float64((spaces + 1))

	return int(math.Round(0.0588*L - 0.296*S - 15.8))
}

func main() {
	text := lib50.GetString("Text: ")
	score := readability(text)

	if score > 16 {
		fmt.Println("Grade 16+")
	} else if score < 1 {
		fmt.Println("Before Grade 1")
	} else {
		fmt.Println("Grade", score)
	}
}
