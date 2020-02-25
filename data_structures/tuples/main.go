package main

import "fmt"

func withLength(input string) (string, int) {
	return input, len(input)
}

func main() {
	example := "Hello"

	word, length := withLength(example)

	fmt.Printf("Word: %s Length: %v", word, length)
}
