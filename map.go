package piscine

func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, len(a))
	for i, r := range a {
		b[i] = f(r)
	}
	return b
}
