package main

import (
	"fmt"
	"testing"
)

func Test_readability(t *testing.T) {
	tests := map[string]struct {
		message string
		input   string
		output  int
	}{
		"single_sentence": {
			input:   "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.",
			message: "handles single sentence with multiple words",
			output:  7,
		},
		"single_sentence_other_punctuation": {
			input:   "There are more things in Heaven and Earth, Horatio, than are dreamt of in your philosophy.",
			message: "handles punctuation within a single sentence",
			output:  9,
		},
		"handles punctuation within a single sentence": {
			input:   "Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do: once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, \"and what is the use of a book,\" thought Alice \"without pictures or conversation?\"",
			message: "handles more complex single sentence",
			output:  8,
		},
		"multiple_sentences": {
			input:   "Harry Potter was a highly unusual boy in many ways. For one thing, he hated the summer holidays more than any other time of year. For another, he really wanted to do his homework, but was forced to do it in secret, in the dead of the night. And he also happened to be a wizard.",
			message: "handles multiple sentences",
			output:  5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if readability(test.input) == test.output {
				fmt.Println(test.message)
			} else {
				t.Error(test.message)
			}
		})
	}
}
