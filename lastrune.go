package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	a := []rune(s)
	return a[len(a)-1]
}
