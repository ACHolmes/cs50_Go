package main

import (
	"fmt"
	"strings"

	"example.com/lib50"
)

type Player int

const (
	PLAYER1 = 1
	PLAYER2 = 2
	TIE     = 0
)

func wordScore(word string) int {
	scores := map[rune]int{
		'a': 1,
		'b': 3,
		'c': 3,
		'd': 2,
		'e': 1,
		'f': 4,
		'g': 2,
		'h': 4,
		'i': 1,
		'j': 8,
		'k': 5,
		'l': 1,
		'm': 3,
		'n': 1,
		'o': 1,
		'p': 3,
		'q': 10,
		'r': 1,
		's': 1,
		't': 1,
		'u': 1,
		'v': 4,
		'w': 4,
		'x': 8,
		'y': 4,
		'z': 10,
	}
	var total int
	for _, c := range word {
		total += scores[c]
	}
	return total
}

func scrabble(p1Input, p2Input string) Player {
	p1Word := strings.ToLower(p1Input)
	p2Word := strings.ToLower(p2Input)
	p1Score, p2Score := wordScore(p1Word), wordScore(p2Word)
	if p1Score > p2Score {
		return PLAYER1
	}
	if p2Score > p1Score {
		return PLAYER2
	}
	return TIE
}

func main() {
	p1Input := lib50.GetString("Player 1: ")
	p2Input := lib50.GetString("Player 2: ")
	switch scrabble(p1Input, p2Input) {
	case PLAYER1:
		fmt.Println("Player 1 wins!")
	case PLAYER2:
		fmt.Println("Player 2 wins!")
	case TIE:
		fmt.Println("Tie!")
	}
}
