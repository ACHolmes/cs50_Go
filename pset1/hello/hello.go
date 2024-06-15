package main

import (
	"fmt"

	"example.com/lib50"
)

func main() {

	name := lib50.GetString("Please enter your name: ")

	fmt.Printf("hello, %s!\n", name)
}
