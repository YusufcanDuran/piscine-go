package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]
	programName := getBaseName(programPath)

	for _, char := range programName {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func getBaseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}
