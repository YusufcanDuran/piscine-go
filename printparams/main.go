package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Program adını dahil etme, sadece argümanları al

	for _, arg := range args {
		printArgument(arg)
	}
}

func printArgument(arg string) {
	for _, char := range arg {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
