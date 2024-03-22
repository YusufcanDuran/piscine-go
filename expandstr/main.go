package main

import (
	"os"

	"github.com/01-edu/z01"
)

func displayStringWithSpaces(input string) {
	// Split the input string into words
	words := splitWords(input)

	// Check if there are any words
	if len(words) == 0 {
		return
	}

	// Display the first word without leading spaces
	firstWord := words[0]
	for _, char := range firstWord {
		z01.PrintRune(char)
	}

	// Display the remaining words with three spaces between each word
	for _, word := range words[1:] {
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		z01.PrintRune(' ')
		for _, char := range word {
			z01.PrintRune(char)
		}
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}

func splitWords(input string) []string {
	// Implement a simple word splitter without using strings.Fields
	var words []string
	var currentWord string
	inWord := false

	for _, char := range input {
		if char != ' ' {
			// Append character to the current word
			currentWord += string(char)
			inWord = true
		} else {
			// Check if we are transitioning from a word to a space
			if inWord {
				words = append(words, currentWord)
				currentWord = ""
				inWord = false
			}
		}
	}

	// Append the last word if it exists
	if inWord {
		words = append(words, currentWord)
	}

	return words
}

func main() {
	// Check if the program is called with exactly one argument
	if len(os.Args) != 2 {
		return
	}

	// Get the input string from the command line argument
	inputString := os.Args[1]

	// Display the string with exactly three spaces between each word
	displayStringWithSpaces(inputString)
}
