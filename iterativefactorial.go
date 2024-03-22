package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 25 {
		return 0
	} else {
		result := 1
		for i := nb; i >= 1; i-- {
			result *= i
		}
		return result
	}
}
