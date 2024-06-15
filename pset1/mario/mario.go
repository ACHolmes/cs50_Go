package main

import (
	"fmt"

	"example.com/lib50"
)

func main() {
	var height int

	for {
		height = lib50.GetInt("Please enter here: ")
		if height > 0 && height <= 8 {
			break
		}
	}

	for row := 0; row < height; row++ {
		for spaces := height - (row + 1); spaces > 0; spaces-- {
			fmt.Print(" ")
		}
		for hashes := 0; hashes < row+1; hashes++ {
			fmt.Print("#")
		}

		fmt.Printf("  ")

		for hashes := 0; hashes < row+1; hashes++ {
			fmt.Print("#")
		}

		fmt.Println()
	}

}
