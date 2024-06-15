package main

import (
	"fmt"
	"testing"
)

func TestScrabble(t *testing.T) {
	tests := map[string]struct {
		message string
		p1Input string
		p2Input string
		output  Player
	}{
		"tie_letter_case": {
			message: "handles letter cases correctly",
			p1Input: "LETTERCASE",
			p2Input: "lettercase",
			output:  TIE,
		},
		"tie_punctuation": {
			message: "handles punctuation correctly",
			p1Input: "Punctuation!?!?",
			p2Input: "punctuation",
			output:  TIE,
		},
		"test1": {
			message: "correctly identifies 'Question?' and 'Question!' as a tie",
			p1Input: "Question?",
			p2Input: "Question!",
			output:  TIE,
		},
		"test2": {
			message: "correctly identifies 'drawing' and 'illustration' as a tie",
			p1Input: "drawing",
			p2Input: "illustration",
			output:  TIE,
		},
		"test3": {
			message: "correctly identifies 'hai!' as winner over 'Oh,'",
			p1Input: "hai!",
			p2Input: "Oh,",
			output:  PLAYER1,
		},
		"test4": {
			message: "correctly identifies 'COMPUTER' as winner over 'science'",
			p1Input: "COMPUTER",
			p2Input: "science",
			output:  PLAYER1,
		},
		"test5": {
			message: "correctly identifies 'Scrabble' as winner over 'wiNNeR'",
			p1Input: "Scrabble",
			p2Input: "wiNNeR",
			output:  PLAYER1,
		},
		"test6": {
			message: "correctly identifies 'pig' as winner over 'dog''",
			p1Input: "pig",
			p2Input: "dog",
			output:  PLAYER1,
		},
		"complex_case": {
			message: "correctly identifies 'Skating!' as winner over 'figure?'",
			p1Input: "figure?",
			p2Input: "Skating!",
			output:  PLAYER2,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if scrabble(test.p1Input, test.p2Input) == test.output {
				fmt.Println(test.message)
			} else {
				t.Errorf(test.message)
			}
		})
	}
}
