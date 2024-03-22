package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	rot13 := ""
	for _, ch := range os.Args[1] {
		if ch >= 'a' && ch <= 'z' {
			rotated := ch + 13
			if rotated > 'z' {
				rotated -= 26
			}
			rot13 += string(rune(rotated))
		} else if ch >= 'A' && ch <= 'Z' {

			rotated := ch + 13
			if rotated > 'Z' {
				rotated -= 26
			}
			rot13 += string(rune(rotated))

		} else {
			rot13 += string(ch)
		}
	}
	for _, ch := range rot13 {
		z01.PrintRune(ch)
	}

	z01.PrintRune('\n')
}
