package main

import (
	"fmt"
	"os"
	"unicode"

	"example.com/lib50"
)

func validateKey(key string) (bool, map[rune]rune) {
	if len(key) != 26 {
		return false, nil
	}
	converter := make(map[rune]rune)
	for i, c := range key {
		converter[rune('a'+i)] = unicode.ToLower(c)
	}
	if len(converter) != 26 {
		return false, nil
	}
	return true, converter
}

func substitution(converter map[rune]rune, plaintext string) string {
	ciphertext := make([]rune, len(plaintext))
	for i, c := range plaintext {
		upper := unicode.IsUpper(c)
		val, ok := converter[unicode.ToLower(c)]
		if ok {
			if upper {
				ciphertext[i] = unicode.ToUpper(val)
			} else {

				ciphertext[i] = val
			}
		} else {
			ciphertext[i] = c
		}
	}
	return string(ciphertext)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./substitution key")
		return
	}

	ok, converter := validateKey(os.Args[1])

	if !ok {
		fmt.Println("Invalid key")
		return
	}

	plaintext := lib50.GetString("plaintext: ")
	fmt.Println("ciphertext:", substitution(converter, plaintext))
}
