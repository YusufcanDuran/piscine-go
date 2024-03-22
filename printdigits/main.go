package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "0123456789"

	for _, char := range alphabet {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
