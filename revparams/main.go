package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for i := len(args) - 1; i >= 0; i-- {
		printArgument(args[i])
	}
}

func printArgument(arg string) {
	for _, char := range arg {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
